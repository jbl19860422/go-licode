package nice

const NICE_CANDIDATE_MAX_FOUNDATION  = (32+1)
/**
 * NiceCandidateType:
 * @NICE_CANDIDATE_TYPE_HOST: A host candidate
 * @NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE: A server reflexive candidate
 * @NICE_CANDIDATE_TYPE_PEER_REFLEXIVE: A peer reflexive candidate
 * @NICE_CANDIDATE_TYPE_RELAYED: A relay candidate
 *
 * An enum represneting the type of a candidate
 */
type NiceCandidateType int
const (
	_ NiceCandidateType = iota
	NICE_CANDIDATE_TYPE_HOST
	NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE
	NICE_CANDIDATE_TYPE_PEER_REFLEXIVE
	NICE_CANDIDATE_TYPE_RELAYED
)

/**
 * NiceCandidateTransport:
 * @NICE_CANDIDATE_TRANSPORT_UDP: UDP transport
 * @NICE_CANDIDATE_TRANSPORT_TCP_ACTIVE: TCP Active transport
 * @NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE: TCP Passive transport
 * @NICE_CANDIDATE_TRANSPORT_TCP_SO: TCP Simultaneous-Open transport
 *
 * An enum representing the type of transport to use
 */
type NiceCandidateTransport int
const (
	_ NiceCandidateTransport = iota
	NICE_CANDIDATE_TRANSPORT_UDP
	NICE_CANDIDATE_TRANSPORT_TCP_ACTIVE
	NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE
	NICE_CANDIDATE_TRANSPORT_TCP_SO
)

/**
 * NiceRelayType:
 * @NICE_RELAY_TYPE_TURN_UDP: A TURN relay using UDP
 * @NICE_RELAY_TYPE_TURN_TCP: A TURN relay using TCP
 * @NICE_RELAY_TYPE_TURN_TLS: A TURN relay using TLS over TCP
 *
 * An enum representing the type of relay to use
*/
type NiceRelayType int
const (
	_ NiceRelayType = iota
	NICE_RELAY_TYPE_TURN_UDP
	NICE_RELAY_TYPE_TURN_TCP
	NICE_RELAY_TYPE_TURN_TLS
)

/* Constants for determining candidate priorities */
const NICE_CANDIDATE_TYPE_PREF_HOST = 120
const NICE_CANDIDATE_TYPE_PREF_PEER_REFLEXIVE = 110
const NICE_CANDIDATE_TYPE_PREF_NAT_ASSISTED = 105
const NICE_CANDIDATE_TYPE_PREF_SERVER_REFLEXIVE = 100
const NICE_CANDIDATE_TYPE_PREF_RELAYED_UDP = 30
const NICE_CANDIDATE_TYPE_PREF_RELAYED = 20
/* Priority preference constants for MS-ICE compatibility */
const NICE_CANDIDATE_TRANSPORT_MS_PREF_UDP = 15
const NICE_CANDIDATE_TRANSPORT_MS_PREF_TCP = 6
const NICE_CANDIDATE_DIRECTION_MS_PREF_PASSIVE = 2
const NICE_CANDIDATE_DIRECTION_MS_PREF_ACTIVE = 5

/**
 * TurnServer:
 * @ref_count: Reference count for the structure.
 * @server: The #NiceAddress of the TURN server
 * @username: The TURN username
 * @password: The TURN password
 * @type: The #NiceRelayType of the server
 *
 * A structure to store the TURN relay settings
 */
type TurnServer struct {
	server 			NiceAddress
	username		string
	password		string
	typ 			NiceRelayType
}

type NiceCandidate struct {
	typ 			NiceCandidateType
	transport 		NiceCandidateTransport
	addr 			NiceAddress
	base_addr		NiceAddress
	priority		uint32
	stream_id 		uint
	component_id 	uint
	foundation		[]byte
	username		string
	password		string
	server			NiceAddress
	turn 			*TurnServer
	sockptr			NiceSockInterface
}

/**
 * nice_candidate_new:
 * @type: The #NiceCandidateType of the candidate to create
 *
 * Creates a new candidate. Must be freed with nice_candidate_free()
 *
 * Returns: A new #NiceCandidate
*/
func nice_candidate_new(typ NiceCandidateType) *NiceCandidate {
	return &NiceCandidate{
		typ:typ,
	}
}

func nice_candidate_copy(cand *NiceCandidate) *NiceCandidate {
	if cand == nil {
		return nil
	}
	c := nice_candidate_new(cand.typ)
	c.transport = cand.transport
	c.addr = cand.addr
	c.base_addr = cand.base_addr
	c.priority = cand.priority
	c.stream_id = cand.stream_id
	c.component_id = cand.component_id
	c.foundation = make([]byte, len(c.foundation))
	copy(c.foundation, cand.foundation)
	c.username = cand.username
	c.password = cand.password
	c.turn = nil
	c.sockptr = cand.sockptr
	return c
}

func nice_candidate_equal(c1 *NiceCandidate, c2 *NiceCandidate) bool {
	if c1.transport == c2.transport && c1.addr == c2.addr {
		return true
	}
	return false
}

func nice_candidate_jingle_priority (candidate *NiceCandidate) uint32 {
	switch candidate.typ {
	case NICE_CANDIDATE_TYPE_HOST:
		return 1000
	case NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE:
		return 900
	case NICE_CANDIDATE_TYPE_PEER_REFLEXIVE:
		return 900
	case NICE_CANDIDATE_TYPE_RELAYED:
		return 500
	default:
		return 0
	}
}

func nice_candidate_msn_priority(candidate *NiceCandidate) uint32 {
	switch candidate.typ {
	case NICE_CANDIDATE_TYPE_HOST:
		return 830
	case NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE:
		return 550
	case NICE_CANDIDATE_TYPE_PEER_REFLEXIVE:
		return 550
	case NICE_CANDIDATE_TYPE_RELAYED:
		return 450
	default:
		return 0
	}
}

func nice_candidate_ice_type_preference (candidate *NiceCandidate, reliable bool, nat_assisted bool) uint8 {
	var type_preference uint8

	switch candidate.typ {
	case NICE_CANDIDATE_TYPE_HOST:
		type_preference = NICE_CANDIDATE_TYPE_PREF_HOST
	case NICE_CANDIDATE_TYPE_PEER_REFLEXIVE:
		type_preference = NICE_CANDIDATE_TYPE_PREF_PEER_REFLEXIVE
	case NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE:
		if nat_assisted {
			type_preference = NICE_CANDIDATE_TYPE_PREF_NAT_ASSISTED
		} else {
			type_preference = NICE_CANDIDATE_TYPE_PREF_SERVER_REFLEXIVE
		}
	case NICE_CANDIDATE_TYPE_RELAYED:
		if candidate.turn.typ == NICE_RELAY_TYPE_TURN_UDP {
			type_preference = NICE_CANDIDATE_TYPE_PREF_RELAYED_UDP
		} else {
			type_preference = NICE_CANDIDATE_TYPE_PREF_RELAYED
		}
	default:
		type_preference = 0
	}

	if ((reliable && candidate.transport == NICE_CANDIDATE_TRANSPORT_UDP) ||
		(!reliable && candidate.transport != NICE_CANDIDATE_TRANSPORT_UDP)) {
		type_preference = type_preference / 2
	}

	return type_preference
}

func nice_candidate_ice_priority (candidate *NiceCandidate, reliable bool, nat_assisted bool) uint32 {
	var type_preference uint8
	var local_preference uint16
	type_preference = nice_candidate_ice_type_preference (candidate, reliable, nat_assisted)
	local_preference = nice_candidate_ice_local_preference (candidate)
	return nice_candidate_ice_priority_full (type_preference, local_preference, candidate.component_id)
}

func nice_candidate_ice_local_preference(candidate *NiceCandidate) uint16 {
	var direction_preference uint
	switch candidate.transport {
	case NICE_CANDIDATE_TRANSPORT_TCP_ACTIVE:
		if candidate.typ == NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE || candidate.typ == NICE_CANDIDATE_TYPE_HOST {
			direction_preference = 4
		} else {
			direction_preference = 6
		}
	case NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE:
		if candidate.typ == NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE || candidate.typ == NICE_CANDIDATE_TYPE_HOST {
			direction_preference = 2
		} else {
			direction_preference = 4
		}
	case NICE_CANDIDATE_TRANSPORT_TCP_SO:
		if candidate.typ == NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE || candidate.typ == NICE_CANDIDATE_TYPE_HOST {
			direction_preference = 6
		} else {
			direction_preference = 2
		}
	case NICE_CANDIDATE_TRANSPORT_UDP:
		return 1
	}
	return uint16(nice_candidate_ice_local_preference_full (direction_preference, uint(nice_candidate_ip_local_preference (candidate))))
}
func nice_candidate_ice_local_preference_full (direction_preference uint, other_preference uint) uint32 {
	return uint32(0x2000 * direction_preference + other_preference)
}

func nice_candidate_ip_local_preference (candidate *NiceCandidate) uint8 {
	var preference uint8 = 0
	var ip_string string
	if candidate.typ == NICE_CANDIDATE_TYPE_HOST {
		ip_string = candidate.addr.ip
	} else {
		ip_string = candidate.base_addr.ip
	}

	addrs, err := nice_interfaces_get_local_ips()
	if err != nil {
		return 0
	}

	for i := 0; i < len(addrs); i++ {
		if ip_string != addrs[i].ip {
			preference++
			continue
		}
		break
	}

	return preference
}

func nice_candidate_ms_ice_local_preference_full(transport_preference uint8, direction_preference uint8, other_preference uint8) uint32 {
	return uint32(0x1000 * uint(transport_preference) + 0x200 * uint(direction_preference) + 0x1 * uint(other_preference))
}


func nice_candidate_ms_ice_local_preference (candidate *NiceCandidate) uint32 {
	var transport_preference uint8 = 0
	var direction_preference uint8 = 0

	switch candidate.transport {
		case NICE_CANDIDATE_TRANSPORT_TCP_SO, NICE_CANDIDATE_TRANSPORT_TCP_ACTIVE:
			transport_preference = NICE_CANDIDATE_TRANSPORT_MS_PREF_TCP
			direction_preference = NICE_CANDIDATE_DIRECTION_MS_PREF_ACTIVE
		case NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE:
			transport_preference = NICE_CANDIDATE_TRANSPORT_MS_PREF_TCP
			direction_preference = NICE_CANDIDATE_DIRECTION_MS_PREF_PASSIVE
		case NICE_CANDIDATE_TRANSPORT_UDP:
			transport_preference = NICE_CANDIDATE_TRANSPORT_MS_PREF_UDP
		default:
			transport_preference = NICE_CANDIDATE_TRANSPORT_MS_PREF_UDP
	}

	return nice_candidate_ms_ice_local_preference_full(transport_preference,
		direction_preference, nice_candidate_ip_local_preference (candidate));
}

func nice_candidate_ms_ice_priority (candidate *NiceCandidate, reliable bool, nat_assisted bool) uint32 {
	var type_preference uint8
	var local_preference uint16
	type_preference = nice_candidate_ice_type_preference(candidate, reliable, nat_assisted)
	local_preference = uint16(nice_candidate_ms_ice_local_preference (candidate))
	return nice_candidate_ice_priority_full(type_preference, local_preference, candidate.component_id)
}

/*
 * ICE 4.1.2.1. "Recommended Formula" (ID-19):
 * returns number between 1 and 0x7effffff
*/
func nice_candidate_ice_priority_full(type_preference uint8, local_preference uint16, component_id	uint) uint32 {
	return uint32(0x1000000 * uint(type_preference) + 0x100 * uint(local_preference) + (0x100 - component_id))
}

/*
 * Assings a foundation to the candidate.
 *
 * Implements the mechanism described in ICE sect
 * 4.1.1.3 "Computing Foundations" (ID-19).
 */
func priv_assign_foundation (agent *NiceAgent, candidate *NiceCandidate) {
	for i := 0; i < len(agent.streams); i++ {
		stream := agent.streams[i]
		for j := 0; j < len(stream.components); j++ {
			component := stream.components[j]
			for k := 0; k < len(component.local_candidates); k++ {
				n := component.local_candidates[k]
				if candidate.typ == n.typ && candidate.transport == n.transport &&
					nice_address_equal_no_port(candidate.base_addr, n.base_addr) &&
					(candidate.typ != NICE_CANDIDATE_TYPE_RELAYED || priv_compare_turn_servers(candidate.turn, n.turn)) &&
					!(agent.compatibility == NICE_COMPATIBILITY_GOOGLE && n.typ == NICE_CANDIDATE_TYPE_RELAYED) {
					candidate.foundation = n.foundation
					if n.username != "" {
						candidate.username = n.username
					}

					if n.password != "" {
						candidate.password = n.password
					}
				}
			}
		}
	}
}

/*
 * Calculates the pair priority as specified in ICE
 * sect 5.7.2. "Computing Pair Priority and Ordering Pairs" (ID-19).
 */
func nice_candidate_pair_priority (o_prio uint32, a_prio uint32) uint64 {
	max := MaxUInt32(o_prio, a_prio)
	min := MinUInt32(o_prio, a_prio)
	var one uint32 = 0
	if o_prio > a_prio {
		one = 1
	}

	var o uint64 = 1
	var tw uint64 = 32
	return uint64(uint32(o << tw) * min + 2 * max + one)
}

