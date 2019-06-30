package nice

import "sync"

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
	local_addresses				[]NiceAddress /* list of NiceAddresses for local interfaces */
	streams						[]NiceStream
	next_candidate_id			uint
	next_stream_id				uint
	//NiceRNG *rng;

}
