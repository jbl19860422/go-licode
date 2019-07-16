package nice

import (
	"time"
	"fmt"
)

const NICE_CANDIDATE_PAIR_MAX_FOUNDATION = NICE_CANDIDATE_MAX_FOUNDATION*2

/**
 * NiceCheckState:
 * @NICE_CHECK_WAITING: Waiting to be scheduled.
 * @NICE_CHECK_IN_PROGRESS: Connection checks started.
 * @NICE_CHECK_SUCCEEDED: Connection successfully checked.
 * @NICE_CHECK_FAILED: No connectivity; retransmissions ceased.
 * @NICE_CHECK_FROZEN: Waiting to be scheduled to %NICE_CHECK_WAITING.
 * @NICE_CHECK_DISCOVERED: A valid candidate pair not on the check list.
 *
 * States for checking a candidate pair.
 */
type NiceCheckState = int
const (
	_ NiceCheckState 	= 	iota
	NICE_CHECK_WAITING 	= 	1
	NICE_CHECK_IN_PROGRESS
	NICE_CHECK_SUCCEEDED
	NICE_CHECK_FAILED
	NICE_CHECK_FROZEN
	NICE_CHECK_DISCOVERED
)

type StunTransaction struct {
	next_tick	time.Time	//GTimeVal next_tick;       /* next tick timestamp */
	timer 		StunTimer
	buffer		[STUN_MAX_MESSAGE_SIZE_IPV6]byte
	message 	*StunMessage
}

type CandidateCheckPair struct {
	stream_id 		uint
	component_id 	uint
	local 			*NiceCandidate
	remote 			*NiceCandidate
	sockptr			NiceSockInterface
	foundation 		[]byte
	state 			NiceCheckState
	nominated		bool
	valid 			bool
	use_candidate_on_next_check	bool
	mark_nominated_on_response_arrival	bool
	retransmit		bool
	discovered_pair *CandidateCheckPair
	succeeded_pair	*CandidateCheckPair
	priority		uint64
	prflx_priority	uint32
	stun_transactions	[]*StunTransaction
}

func NewCandidateCheckPair() *CandidateCheckPair {
	return &CandidateCheckPair{}
}

func ensure_unique_priority(stream *NiceStream, component *NiceComponent, priority uint32) uint32 {
	if priority == 0 {
		priority--
	}
again:
	for i := 0; i < len(component.local_candidates); i++ {
		if component.local_candidates[i].priority == priority {
			priority--
			goto again
		}
	}

	for i := 0; i < len(stream.conncheck_list); i++ {
		p := stream.conncheck_list[i]
		if p.component_id == component.id && p.prflx_priority == priority {
			priority--
			goto again
		}
	}
	return priority
}

/*
 * Forms new candidate pairs by matching the new local candidate
 * 'local_cand' with all existing remote candidates of 'component'.
 *
 * @param agent context
 * @param component pointer to the component
 * @param local local candidate to match with
 *
 * @return number of checks added, negative on fatal errors
 */
func conn_check_add_for_local_candidate (agent *NiceAgent, stream_id uint, component *NiceComponent, local *NiceCandidate) int {
/*
 * note: according to 7.1.3.2.1 "Discovering Peer Reflexive
 * Candidates", the peer reflexive candidate is not paired
 * with other remote candidates
 */
 	var added int = 0
 	if agent.compatibility == NICE_COMPATIBILITY_RFC5245 && local.typ == NICE_CANDIDATE_TYPE_PEER_REFLEXIVE {
 		return 0	//todo
	}

	for i := 0; i < len(component.remote_candidates); i++ {
		remote := component.remote_candidates[i]
		ret := conn_check_add_for_candidate_pair (agent, stream_id, component, local, remote)
		if ret {
			added++
		}
	}
 	return added
}

func conn_check_match_transport(transport NiceCandidateTransport) NiceCandidateTransport {
	switch transport {
		case NICE_CANDIDATE_TRANSPORT_TCP_ACTIVE:
			return NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE
		case NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE:
			return NICE_CANDIDATE_TRANSPORT_TCP_ACTIVE
		case NICE_CANDIDATE_TRANSPORT_TCP_SO, NICE_CANDIDATE_TRANSPORT_UDP:
			return transport
		default:
			return transport
	}
}

func conn_check_add_for_candidate_pair(agent *NiceAgent, stream_id uint, component *NiceComponent, local *NiceCandidate, remote *NiceCandidate) bool {
	var ret bool = false
	/* note: do not create pairs where the local candidate is
 *       a srv-reflexive (ICE 5.7.3. "Pruning the pairs" ID-9) */
	if (agent.compatibility == NICE_COMPATIBILITY_RFC5245 || agent.compatibility == NICE_COMPATIBILITY_WLM2009 ||
		agent.compatibility == NICE_COMPATIBILITY_OC2007R2) && local.typ == NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE {
		return false
	}
	/* note: do not create pairs where local candidate has TCP passive transport
 *       (ice-tcp-13 6.2. "Forming the Check Lists") */
 	if local.transport == NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE {
 		return false
	}

	if local.transport == conn_check_match_transport(remote.transport) && EqualFamily(local.addr, remote.addr) {

	}
	return ret
}

func priv_conn_check_add_for_candidate_pair_matched(agent *NiceAgent, stream_id uint, component *NiceComponent, local *NiceCandidate, remote *NiceCandidate, initial_state NiceCheckState) *CandidateCheckPair {
	var pair *CandidateCheckPair

	pair = priv_add_new_check_pair (agent, stream_id, component, local, remote, initial_state)
	if component.state == NICE_COMPONENT_STATE_CONNECTED || component.state == NICE_COMPONENT_STATE_READY {
		agent.agent_signal_component_state_change(stream_id, component.id, NICE_COMPONENT_STATE_CONNECTED)
	} else {
		agent.agent_signal_component_state_change(stream_id, component.id, NICE_COMPONENT_STATE_CONNECTING)
	}

	return pair
}

func InsertSorted(list []*CandidateCheckPair, pair *CandidateCheckPair) []*CandidateCheckPair {
	var i int = 0
	for i = 0; i < len(list); i++ {
		if pair.priority >= list[i].priority {
			break
		}
	}

	rear := append([]*CandidateCheckPair{}, list[i:]...)
	return append(append(list[:i], pair), rear...)
}
/*
 * Creates a new connectivity check pair and adds it to
 * the agent's list of checks.
 */
func priv_add_new_check_pair(agent *NiceAgent,stream_id uint, component *NiceComponent, local *NiceCandidate, remote *NiceCandidate, initial_state NiceCheckState) *CandidateCheckPair {
	var stream *NiceStream
	var pair *CandidateCheckPair

	stream = agent.find_stream(stream_id)
	pair = NewCandidateCheckPair()

	pair.stream_id = stream_id
	pair.component_id = component.id
	pair.local = local
	pair.remote = remote
	/* note: we use the remote sockptr only in the case
	 * of TCP transport
	*/
	if local.transport ==  NICE_CANDIDATE_TRANSPORT_TCP_PASSIVE && remote.typ == NICE_CANDIDATE_TYPE_PEER_REFLEXIVE {
		pair.sockptr = remote.sockptr
	} else {
		pair.sockptr = local.sockptr
	}
	pair.foundation = []byte(string(local.foundation) + ":" + string(remote.foundation))
	pair.priority = agent.agent_candidate_pair_priority(local, remote)

	pair.state = initial_state
	fmt.Println("agent creating a new pair")
	pair.prflx_priority = ensure_unique_prflx_priority (stream, component, local.priority, peer_reflexive_candidate_priority (agent, local))

	stream.conncheck_list = InsertSorted(stream.conncheck_list, pair)

	/* implement the hard upper limit for number of
	   checks (see sect 5.7.3 ICE ID-19): */
	if (agent.compatibility == NICE_COMPATIBILITY_RFC5245) {
		//todo limit conncheck list
		//stream.conncheck_list = priv_limit_conn_check_list_size (agent, stream.conncheck_list, agent.max_conn_checks)
	}

	return pair
}

func peer_reflexive_candidate_priority(agent *NiceAgent, local_candidate *NiceCandidate) uint32  {
	var candidate_priority *NiceCandidate = nice_candidate_new (NICE_CANDIDATE_TYPE_PEER_REFLEXIVE)
	var priority uint32

	candidate_priority.transport = local_candidate.transport
	candidate_priority.component_id = local_candidate.component_id
	candidate_priority.base_addr = local_candidate.addr

	if agent.compatibility == NICE_COMPATIBILITY_GOOGLE {
		priority = nice_candidate_jingle_priority (candidate_priority)
	} else if agent.compatibility == NICE_COMPATIBILITY_MSN || agent.compatibility == NICE_COMPATIBILITY_OC2007 {
		priority = nice_candidate_msn_priority (candidate_priority)
	} else if agent.compatibility == NICE_COMPATIBILITY_OC2007R2 {
		priority = nice_candidate_ms_ice_priority (candidate_priority, agent.reliable, false)
	} else {
		priority = nice_candidate_ice_priority (candidate_priority, agent.reliable, false)
	}

	return priority
}


func ensure_unique_prflx_priority(stream *NiceStream, component *NiceComponent, local_priority uint32, prflx_priority uint32) uint32 {
	/* First, ensure we provide the same value for pairs having
	 * the same local candidate, ie the same local candidate priority
	 * for the sake of coherency with the stun server behaviour that
	 * stores a unique priority value per remote candidate, from the
	 * first stun request it receives (it depends on the kind of NAT
	 * typically, but for NAT that preserves the binding this is required).
	 */
	 for i := 0; i < len(stream.conncheck_list); i++ {
		p := stream.conncheck_list[i]
		if p.component_id == component.id && p.local.priority == local_priority {
			return p.prflx_priority
		}
	 }

	 /* Second, ensure uniqueness across all other prflx_priority values */
again:
	if prflx_priority == 0 {
		prflx_priority--
	}

	for i := 0; i < len(stream.conncheck_list); i++ {
		p := stream.conncheck_list[i]
		if p.component_id == component.id && p.prflx_priority == prflx_priority {
			prflx_priority--
			goto again
		}
	}
	return prflx_priority
}


