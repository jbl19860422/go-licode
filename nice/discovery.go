package nice

type CandidateDiscovery struct {
	typ 			NiceCandidateType
	nicesock		NiceSockInterface
	server 			NiceAddress	/* STUN/TURN server address */
	//GTimeVal next_tick;       /* next tick timestamp */
	pending			bool
	done 			bool
	stream_id		uint
	component_id 	uint
	turn 			*TurnServer

}
