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
