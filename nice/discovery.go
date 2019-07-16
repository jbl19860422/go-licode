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

		
	}
{
static int tick_counter = 0;
if (tick_counter++ % 50 == 0)
nice_debug ("Agent %p : discovery tick #%d with list %p (1)", agent, tick_counter, agent->discovery_list);
}

for (i = agent->discovery_list; i ; i = i->next) {
cand = i->data;

if (cand->pending != TRUE) {
cand->pending = TRUE;

if (agent->discovery_unsched_items)
--agent->discovery_unsched_items;

if (nice_debug_is_enabled ()) {
gchar tmpbuf[INET6_ADDRSTRLEN];
nice_address_to_string (&cand->server, tmpbuf);
nice_debug ("Agent %p : discovery - scheduling cand type %u addr %s.",
agent, cand->type, tmpbuf);
}
if (nice_address_is_valid (&cand->server) &&
(cand->type == NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE ||
cand->type == NICE_CANDIDATE_TYPE_RELAYED)) {
NiceComponent *component;

if (agent_find_component (agent, cand->stream_id,
cand->component_id, NULL, &component) &&
(component->state == NICE_COMPONENT_STATE_DISCONNECTED ||
component->state == NICE_COMPONENT_STATE_FAILED))
agent_signal_component_state_change (agent,
cand->stream_id,
cand->component_id,
NICE_COMPONENT_STATE_GATHERING);

if (cand->type == NICE_CANDIDATE_TYPE_SERVER_REFLEXIVE) {
buffer_len = stun_usage_bind_create (&cand->stun_agent, &cand->stun_message, cand->stun_buffer, sizeof(cand->stun_buffer));
} else if (cand->type == NICE_CANDIDATE_TYPE_RELAYED) {
uint8_t *username = (uint8_t *)cand->turn->username;
gsize username_len = strlen (cand->turn->username);
uint8_t *password = (uint8_t *)cand->turn->password;
gsize password_len = strlen (cand->turn->password);
StunUsageTurnCompatibility turn_compat =
agent_to_turn_compatibility (agent);

if (turn_compat == STUN_USAGE_TURN_COMPATIBILITY_MSN ||
turn_compat == STUN_USAGE_TURN_COMPATIBILITY_OC2007) {
username = cand->turn->decoded_username;
password = cand->turn->decoded_password;
username_len = cand->turn->decoded_username_len;
password_len = cand->turn->decoded_password_len;
}

buffer_len = stun_usage_turn_create (&cand->stun_agent,
&cand->stun_message,  cand->stun_buffer, sizeof(cand->stun_buffer),
cand->stun_resp_msg.buffer == NULL ? NULL : &cand->stun_resp_msg,
STUN_USAGE_TURN_REQUEST_PORT_NORMAL,
-1, -1,
username, username_len,
password, password_len,
turn_compat);
}

if (buffer_len > 0) {
if (nice_socket_is_reliable (cand->nicesock)) {
stun_timer_start_reliable (&cand->timer, agent->stun_reliable_timeout);
} else {
stun_timer_start (&cand->timer,
agent->stun_initial_timeout,
agent->stun_max_retransmissions);
}

/* send the conncheck */
agent_socket_send (cand->nicesock, &cand->server,
buffer_len, (gchar *)cand->stun_buffer);

/* case: success, start waiting for the result */
g_get_current_time (&cand->next_tick);

} else {
/* case: error in starting discovery, start the next discovery */
cand->done = TRUE;
cand->stun_message.buffer = NULL;
cand->stun_message.buffer_len = 0;
continue;
}
}
else
/* allocate relayed candidates */
g_assert_not_reached ();

++not_done; /* note: new discovery scheduled */
}

if (cand->done != TRUE) {
GTimeVal now;

g_get_current_time (&now);

if (cand->stun_message.buffer == NULL) {
nice_debug ("Agent %p : STUN discovery was cancelled, marking discovery done.", agent);
cand->done = TRUE;
}
else if (priv_timer_expired (&cand->next_tick, &now)) {
switch (stun_timer_refresh (&cand->timer)) {
case STUN_USAGE_TIMER_RETURN_TIMEOUT:
{
/* Time out */
/* case: error, abort processing */
StunTransactionId id;

stun_message_id (&cand->stun_message, id);
stun_agent_forget_transaction (&cand->stun_agent, id);

cand->done = TRUE;
cand->stun_message.buffer = NULL;
cand->stun_message.buffer_len = 0;
nice_debug ("Agent %p : bind discovery timed out, aborting discovery item.", agent);
break;
}
case STUN_USAGE_TIMER_RETURN_RETRANSMIT:
{
/* case: not ready complete, so schedule next timeout */
unsigned int timeout = stun_timer_remainder (&cand->timer);

stun_debug ("STUN transaction retransmitted (timeout %dms).",
timeout);

/* retransmit */
agent_socket_send (cand->nicesock, &cand->server,
stun_message_length (&cand->stun_message),
(gchar *)cand->stun_buffer);

/* note: convert from milli to microseconds for g_time_val_add() */
cand->next_tick = now;
g_time_val_add (&cand->next_tick, timeout * 1000);

++not_done; /* note: retry later */
break;
}
case STUN_USAGE_TIMER_RETURN_SUCCESS:
{
unsigned int timeout = stun_timer_remainder (&cand->timer);

cand->next_tick = now;
g_time_val_add (&cand->next_tick, timeout * 1000);

++not_done; /* note: retry later */
break;
}
default:
/* Nothing to do. */
break;
}

} else {
++not_done; /* note: discovery not expired yet */
}
}
}

if (not_done == 0) {
nice_debug ("Agent %p : Candidate gathering FINISHED, stopping discovery timer.", agent);

discovery_free (agent);

agent_gathering_done (agent);

/* note: no pending timers, return FALSE to stop timer */
return FALSE;
}

return TRUE;
}