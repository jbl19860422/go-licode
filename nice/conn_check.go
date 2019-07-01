package nice

import (
	"time"
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
	componet_id 	uint
	local 			*NiceCandidate
	remote 			*NiceCandidate
	sockptr			NiceSockInterface
	foundation 		[NICE_CANDIDATE_PAIR_MAX_FOUNDATION]byte
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
		if p.componet_id == component.id && p.prflx_priority == priority {
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

	if local.transport == 
	return ret
}

/* note: match pairs only if transport and address family are the same */
if (local->transport == conn_check_match_transport (remote->transport) &&
local->addr.s.addr.sa_family == remote->addr.s.addr.sa_family) {
priv_conn_check_add_for_candidate_pair_matched (agent, stream_id, component,
local, remote, NICE_CHECK_FROZEN);
ret = TRUE;
}

return ret;
}

