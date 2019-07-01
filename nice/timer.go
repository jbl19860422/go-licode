package nice

import "time"

/**
 * StunTimer:
 *
 * An opaque structure representing a STUN transaction retransmission timer
*/
type StunTimer struct {
	deadline 			time.Time
	delay 				uint32
	retransmissions 	uint32
	max_retransmissions uint32
}
