package nice

import "encoding/base64"

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

func (this *NiceAgent)discovery_add_local_host_candidate(
										stream_id uint,
										component_id uint,
										address *NiceAddress,
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
	return nil, HOST_CANDIDATE_SUCCESS
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
