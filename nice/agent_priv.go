package nice

import (
	"sync"
	"errors"
	"net"
)

/* XXX: starting from ICE ID-18, Ta SHOULD now be set according
 *      to session bandwidth -> this is not yet implemented in NICE */

const NICE_AGENT_TIMER_TA_DEFAULT = 20      /* timer Ta, msecs (impl. defined) */
const NICE_AGENT_TIMER_TR_DEFAULT = 25000   /* timer Tr, msecs (impl. defined) */
const NICE_AGENT_MAX_CONNECTIVITY_CHECKS_DEFAULT = 100 /* see spec 5.7.3 (ID-19) */

/* An upper limit to size of STUN packets handled (based on Ethernet
 * MTU and estimated typical sizes of ICE STUN packet */
const MAX_STUN_DATAGRAM_PAYLOAD  = 1300

/* maximum number of validates remote candidates to keep, the number is arbitrary but hopefully large enough */
const NICE_COMPONENT_MAX_VALID_CANDIDATES = 50

/* A convenient macro to test if the agent is compatible with RFC5245
 * or OC2007R2. Specifically these two modes share the support
 * of the regular or aggressive nomination mode */
func NICE_AGENT_IS_COMPATIBLE_WITH_RFC5245_OR_OC2007R2(a *NiceAgent) bool {
	return a.compatibility == NICE_COMPATIBILITY_RFC5245 || a.compatibility == NICE_COMPATIBILITY_OC2007R2
}

/**
 * NiceAgentRecvFunc:
 * @agent: The #NiceAgent Object
 * @stream_id: The id of the stream
 * @component_id: The id of the component of the stream
 *        which received the data
 * @len: The length of the data
 * @buf: The buffer containing the data received
 * @user_data: The user data set in nice_agent_attach_recv()
 *
 * Callback function when data is received on a component
 *
*/
type NiceAgentRecvFunc 		func(agent *NiceAgent, stream_id uint, component_id uint,buf []byte, user_data []byte)
type GatheringDoneCb		func(agent *NiceAgent, stream_id uint, data interface{})
type NewSelectPairCb		func(agent *NiceAgent, stream_id uint, component_id uint, foundation []byte, rfoundation []byte, data interface{})
type ComponentStateChangeCb func(agent *NiceAgent, stream_id uint, component_id uint, state uint, data interface{})


type NiceAgent struct {
	agent_mutex 				sync.Mutex
	full_mode					bool
	stun_server_ip 				string
	stun_server_port 			uint16
	proxy_ip					string
	proxy_port					uint16
	proxy_type					NiceProxyType
	proxy_username				string
	proxy_password				string
	saved_controlling_mode 		bool
	timer_ta					uint
	max_conn_checks				uint
	force_relay					bool
	stun_max_retransmissions 	uint
	stun_initial_timeout		uint
	stun_reliable_timeout		uint
	nomination_mode				NiceNominationMode
	support_renomination		bool
	local_addresses				[]net.Addr
	streams						[]*NiceStream
	next_candidate_id			uint
	next_stream_id				uint
	rng 						*NiceRNG
	discovery_list				[]*CandidateDiscovery
	use_ice_trickle				bool

	compatibility				NiceCompatibility	/* property: Compatibility mode */
	use_ice_udp					bool
	use_ice_tcp					bool

	software_attribute 			string       /* SOFTWARE attribute */
	reliable					bool         /* property: reliable */
	keepalive_conncheck			bool

	stun_addr					string
	stun_port					uint16
	controlling_mode			bool

	gathering_done_db			GatheringDoneCb
	new_selectpair_cb			NewSelectPairCb
	componet_state_change_cb	ComponentStateChangeCb
}

func NewNiceAgent() *NiceAgent {
	a := &NiceAgent{}
	a.rng = NewNiceRNG()
	return a
}

func (this *NiceAgent) SetStunServer(addr string) {
	this.stun_addr = addr
}

func (this *NiceAgent) SetStunPort(port uint16) {
	this.stun_port = port
}

func (this *NiceAgent) SetControllingMode(mode bool) {
	this.controlling_mode = mode
}

func (this *NiceAgent) nice_agent_add_stream(n_components uint) uint {
	if n_components <= 0 {
		return 0
	}

	this.agent_mutex.Lock()
	defer this.agent_mutex.Unlock()
	stream := NewNiceStream(this.next_stream_id, n_components, this)

	this.streams = append(this.streams, stream)
	if this.reliable {
		var i uint = 0
		for i = 0; i < n_components; i++ {
			c := stream.find_component_by_id(i + 1)
			if c != nil {
				//todo add pseudo tcp, but first we don't use pseudo tcp
				//pseudo_tcp_socket_create (agent, stream, component);
			}
		}
	}
	//todo optimize rng, we not need to put it in agent, just move it to utils
	stream.nice_stream_initialize_credentials(this.rng)
	//todo need to agent_unlock_and_emit
	return stream.id
}

func (this *NiceAgent) agent_find_component(stream_id uint, component_id uint) (*NiceStream, *NiceComponent) {
	stream := this.find_stream(stream_id)
	if stream == nil {
		return nil, nil
	}

	component := stream.find_component_by_id(component_id)
	return stream, component
}

func (this *NiceAgent) find_stream(stream_id uint) *NiceStream {
	for i := 0; i < len(this.streams); i++ {
		if this.streams[i].id == stream_id {
			return this.streams[i]
		}
	}
	return nil
}

func (this *NiceAgent) nice_agent_attach_recv(stream_id uint, component_id uint, recv_func NiceAgentRecvFunc, data interface{}) bool {
	if stream_id < 1 {
		return false
	}

	if component_id < 1 {
		return false
	}

	this.agent_mutex.Lock()
	defer this.agent_mutex.Unlock()

	s, c := this.agent_find_component(stream_id, component_id)
	if s == nil || c == nil {
		return false
	}

	c.nice_component_set_io_callback(recv_func, nil, nil)
	return true
}

func (this *NiceAgent) nice_agent_set_port_range(stream_id uint, component_id uint, min_port uint, max_port uint) {
	this.agent_mutex.Lock()
	defer this.agent_mutex.Unlock()

	_, c := this.agent_find_component(stream_id, component_id)
	if c == nil {
		return
	}
	c.min_port = min_port
	c.max_port = max_port
}

const ADD_HOST_MIN = 0
const ADD_HOST_UDP = ADD_HOST_MIN
const ADD_HOST_TCP_ACTIVE = 1
const ADD_HOST_TCP_PASSIVE = 2
const ADD_HOST_MAX = ADD_HOST_TCP_PASSIVE

func (this *NiceAgent) nice_agent_gather_candidates(stream_id uint) error {
	this.agent_mutex.Lock()
	defer this.agent_mutex.Unlock()

	stream := this.find_stream(stream_id)
	if stream == nil {
		return errors.New("could not find the stream")
	}

	if stream.gathering_started {
		return nil
	}

	/* if no local addresses added, generate them ourselves */
	if this.local_addresses == nil {
		var err error
		this.local_addresses, err = nice_interfaces_get_local_ips()
		if err != nil {
			return err
		}
	}

	for cid := 1; cid < len(stream.components); cid++ {
		_, component := this.agent_find_component(stream_id, uint(cid))
		if component == nil {
			continue
		}
		/* generate a local host candidate for each local address */
		for i := 0; i < len(this.local_addresses); i++ {
			for add_type := ADD_HOST_MIN; add_type <= ADD_HOST_MAX; add_type++ {
				var transport NiceCandidateTransport
				var current_port uint
				var start_port	uint
				if this.use_ice_udp == false && add_type == ADD_HOST_UDP {
					continue
				}

				if this.use_ice_tcp == false && add_type != ADD_HOST_UDP {
					continue
				}

				switch add_type {
				case ADD_HOST_UDP:
					transport = NICE_CANDIDATE_TRANSPORT_UDP
				case ADD_HOST_TCP_ACTIVE:
					transport = NICE_CANDIDATE_TRANSPORT_TCP_ACTIVE
				case ADD_HOST_TCP_PASSIVE:
					transport = NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE
				default:
					transport = NICE_CANDIDATE_TRANSPORT_UDP
				}

				start_port = component.min_port
				if component.min_port != 0 {
					start_port = this.rng.rng_generate_int(component.min_port, component.max_port)
				}

				current_port = start_port


			}
		}
	}
}