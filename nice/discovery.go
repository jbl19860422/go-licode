package nice

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
