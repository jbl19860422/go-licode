package nice
const STUN_MESSAGE_TRANS_ID_LEN = 16

/**
 * STUN_AGENT_MAX_SAVED_IDS:
 *
 * Maximum number of simultaneously ongoing STUN transactions.
*/
const STUN_AGENT_MAX_SAVED_IDS 	= 	200

const STUN_MESSAGE_TYPE_LEN 	= 	2

const STUN_MAGIC_COOKIE 		= 0x2112A442
const STUN_MAGIC_COOKIE_LEN		= 4
const STUN_MAX_MESSAGE_SIZE_IPV6 = 1280

const NICE_STREAM_DEF_UFRAG		= 4 + 1
const NICE_STREAM_DEF_PWD     	= 22 + 1   /* pwd + NULL */