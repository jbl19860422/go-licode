package nice

/* A pair of a socket and the GSource which polls it from the main loop. All
 * GSources in a Component must be attached to the same main context:
 * component->ctx.
 *
 * Socket must be non-NULL, but source may be NULL if it has been detached.
 *
 * The Component is stored so this may be used as the user data for a GSource
 * callback. */
type  SocketSource struct {
	socket 		NiceSockInterface
	//GSource *source;//todo
	component 	*NiceComponent
}

type CandidatePairKeepalive struct {
//GSource *tick_source;//todo
	stream_id		uint
	component_id	uint
	timer 			StunTimer
	stun_buffer		[]byte
	stun_message	*StunMessage
}

type CandidatePair struct {
	local			*NiceCandidate
	remote			*NiceCandidate
	priority		uint64
	prflx_priority	uint32
	keepalive		CandidatePairKeepalive
}

type IncomingCheck struct {
	from 			NiceAddress
	local_socket 	NiceSockInterface
	priority		uint32
	use_candidate	bool
	username		string
}

type NiceComponent struct {
	agent 				*NiceAgent
	stream 				*NiceStream
	typ 				NiceComponentType
	id 					uint
	state 				NiceComponentState
	local_candidates	[]*NiceCandidate
	remote_candidates	[]*NiceCandidate
	valid_candidates	[]*NiceCandidate
	socket_sources		*SocketSource
	socket_sources_age	uint
	incoming_checks		[]*IncomingCheck
	turn_servers		[]*TurnServer
	selected_pair		CandidatePair
	io_callback			NiceAgentRecvFunc   /* function called on io cb */

	min_port			uint
	max_port			uint
}

func NewNiceComponent(agent *NiceAgent, stream *NiceStream, id uint) *NiceComponent {
	return &NiceComponent{
		agent:agent,
		stream:stream,
		id:id,
		min_port:0,
		max_port:65535,
	}
}

func (this *NiceComponent) nice_component_set_io_callback(recv_func NiceAgentRecvFunc, user_data interface{}, recv_messages *NiceInputMessage) error {
	this.io_callback = recv_func
	return nil
}
