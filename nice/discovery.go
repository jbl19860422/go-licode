package nice

import (
	"encoding/base64"
	"fmt"
)

type HostCandidateResult int
const (
	_HostCandidateResult = iota
	HOST_CANDIDATE_SUCCESS
	HOST_CANDIDATE_FAILED
	HOST_CANDIDATE_CANT_CREATE_SOCKET
	HOST_CANDIDATE_REDUNDANT
)

type CandidateDiscovery struct {
	typ 				NiceCandidateType
	nicesock			NiceSockInterface
	server 				NiceAddress	/* STUN/TURN server address */
	//GTimeVal next_tick;       /* next tick timestamp */
	pending				bool
	done 				bool
	stream_id			uint
	component_id 		uint
	turn 				*TurnServer
	stun_agent 			StunAgent
	timer 				StunTimer
	stun_buffer			[]byte
	stun_message 		*StunMessage
	stun_resp_buffer 	[]byte
	stun_resp_message	*StunMessage
}

func NewCandidateDiscovery() *CandidateDiscovery {
	return &CandidateDiscovery{}
}

func (this *NiceAgent)discovery_add_local_host_candidate(
										stream_id uint,
										component_id uint,
										address NiceAddress,
										transport NiceCandidateTransport) (*NiceCandidate, HostCandidateResult) {
	s, c := this.agent_find_component(stream_id, component_id)
	if s == nil || c == nil {
		return nil, HOST_CANDIDATE_FAILED
	}

	candidate := nice_candidate_new(NICE_CANDIDATE_TYPE_HOST)
	candidate.transport = transport
	candidate.stream_id = stream_id
	candidate.addr = address
	candidate.base_addr = address

	if this.compatibility == NICE_COMPATIBILITY_GOOGLE {
		candidate.priority = nice_candidate_jingle_priority(candidate)
	} else if this.compatibility == NICE_COMPATIBILITY_MSN || this.compatibility == NICE_COMPATIBILITY_OC2007 {
		candidate.priority = nice_candidate_msn_priority(candidate)
	} else if this.compatibility == NICE_COMPATIBILITY_OC2007R2 {
		candidate.priority = nice_candidate_ms_ice_priority(candidate, this.reliable, false)
	} else {
		candidate.priority = nice_candidate_ice_priority(candidate, this.reliable, false)
	}

	candidate.priority = ensure_unique_priority(s, c, candidate.priority)
	this.priv_generate_candidate_credentials(candidate)
	priv_assign_foundation(this, candidate)

	var nicesock NiceSockInterface
	if transport == NICE_CANDIDATE_TRANSPORT_UDP {
		nicesock = nice_udp_bsd_socket_new(address)
		if nicesock == nil {
			fmt.Println("nice_udp_bsd_socket_new failed")
			return nil, HOST_CANDIDATE_FAILED
		}
	} else {
		return nil, HOST_CANDIDATE_FAILED
	}

	candidate.sockptr = nicesock
	candidate.addr = address
	candidate.base_addr = address
	return candidate, HOST_CANDIDATE_SUCCESS
}

func (this *NiceAgent) priv_generate_candidate_credentials (candidate *NiceCandidate) {
	if (this.compatibility == NICE_COMPATIBILITY_MSN || this.compatibility == NICE_COMPATIBILITY_OC2007) {
		username := this.rng.rng_generate_bytes(32)
		password := this.rng.rng_generate_bytes(16)
		candidate.username = base64.StdEncoding.EncodeToString (username)
		candidate.password = base64.StdEncoding.EncodeToString (password)
	} else if (this.compatibility == NICE_COMPATIBILITY_GOOGLE) {
		candidate.password = ""
		candidate.username = string(this.rng.nice_rng_generate_bytes_print (16))
	}
}

func priv_compare_turn_servers (turn1 *TurnServer, turn2 *TurnServer) bool {
	if turn1 == turn2 {
		return true
	}
	if turn1 == nil || turn2 == nil {
		return false
	}

	return nice_address_equal_no_port (turn1.server, turn2.server)
}

/*
 * Adds a new local candidate. Implements the candidate pruning
 * defined in ICE spec section 4.1.3 "Eliminating Redundant
 * Candidates" (ID-19).
 */
func priv_add_local_candidate_pruned(agent *NiceAgent, stream_id uint, component *NiceComponent, candidate *NiceCandidate) bool {
	for i := 0; i < len(component.local_candidates); i++ {
		c := component.local_candidates[i]
		if nice_address_equal(c.base_addr, candidate.base_addr) && nice_address_equal(c.addr, candidate.addr) && c.transport == candidate.transport {
			return false
		}
	}

	component.local_candidates = append(component.local_candidates, candidate)
	conn_check_add_for_local_candidate(agent, stream_id, component, candidate)
	return true
}

func discovery_schedule(agent *NiceAgent)  {
	if agent.discovery_unsched_items > 0 {
		var res bool = false
		if agent.timer == nil {
			/* step: run first iteration immediately */
			res = priv_discovery_tick_unlocked(agent)
			_ = res
		}
		// todo
		//if res {
		//	agent_timeout_add_with_context (agent, &agent->discovery_timer_source,
		//		"Candidate discovery tick", agent->timer_ta,
		//		priv_discovery_tick_agent_locked, NULL);
		//}
	}
}

/*
 * Timer callback that handles scheduling new candidate discovery
 * processes (paced by the Ta timer), and handles running of the
 * existing discovery processes.
 *
 * This function is designed for the g_timeout_add() interface.
 *
 * @return will return FALSE when no more pending timers.
*/
func priv_discovery_tick_unlocked(agent *NiceAgent) bool {
	var not_done bool = false
	var buffer_len uint32 = 0

	for i := len(agent.discovery_list) - 1; i >= 0; i-- {
		cand := agent.discovery_list[i]
		if !cand.pending {
			cand.pending = true
		}

		if agent.discovery_unsched_items > 0 {
			agent.discovery_unsched_items--
		}

		if cand.typ == NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE || cand.typ == NICE_CANDIDATE_TYPE_RELAYED {
			s, c := agent.agent_find_component(cand.stream_id, cand.component_id)
			if s != nil && c != nil {
				if c.state == NICE_COMPONENT_STATE_DISCONNECTED || c.state == NICE_COMPONENT_STATE_FAILED {
					agent_signal_component_state_change(agent, cand.stream_id, cand.component_id, NICE_COMPONENT_STATE_GATHERING)
				}

				if cand.typ == NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE {

				}
			}
		}

	}
	return false
}