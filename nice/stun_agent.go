package nice


/**
 * StunCompatibility:
 * @STUN_COMPATIBILITY_RFC3489: Use the STUN specifications compatible with
 * RFC 3489
 * @STUN_COMPATIBILITY_RFC5389: Use the STUN specifications compatible with
 * RFC 5389
 * @STUN_COMPATIBILITY_MSICE2: Use the STUN specifications compatible with
 * [MS-ICE2] (a mix between RFC3489 and RFC5389)
 * @STUN_COMPATIBILITY_OC2007: Use the STUN specifications compatible with
 * Microsoft Office Communicator 2007 (basically RFC3489 with swapped
 * REALM and NONCE attribute hex IDs, attributes are not aligned)
 * @STUN_COMPATIBILITY_WLM2009: An alias for @STUN_COMPATIBILITY_MSICE2
 * @STUN_COMPATIBILITY_LAST: Dummy last compatibility mode
 *
 * Enum that specifies the STUN compatibility mode of the #StunAgent
 *
 * <warning>@STUN_COMPATIBILITY_WLM2009 is deprecated and should not be used
 * in newly-written code. It is kept for compatibility reasons and represents
 * the same compatibility as @STUN_COMPATIBILITY_MSICE2.</warning>
*/
type StunCompatibility int
const (
	_ StunCompatibility = iota
	STUN_COMPATIBILITY_RFC3489
	STUN_COMPATIBILITY_RFC5389
	STUN_COMPATIBILITY_MSICE2
	STUN_COMPATIBILITY_OC2007
	STUN_COMPATIBILITY_WLM2009 = STUN_COMPATIBILITY_MSICE2
	STUN_COMPATIBILITY_LAST = STUN_COMPATIBILITY_OC2007
)

/**
 * StunValidationStatus:
 * @STUN_VALIDATION_SUCCESS: The message is validated
 * @STUN_VALIDATION_NOT_STUN: This is not a valid STUN message
 * @STUN_VALIDATION_INCOMPLETE_STUN: The message seems to be valid but incomplete
 * @STUN_VALIDATION_BAD_REQUEST: The message does not have the cookie or the
 * fingerprint while the agent needs it with its usage
 * @STUN_VALIDATION_UNAUTHORIZED_BAD_REQUEST: The message is valid but
 * unauthorized with no username and message-integrity attributes.
 * A BAD_REQUEST error must be generated
 * @STUN_VALIDATION_UNAUTHORIZED: The message is valid but unauthorized as
 * the username/password do not match.
 * An UNAUTHORIZED error must be generated
 * @STUN_VALIDATION_UNMATCHED_RESPONSE: The message is valid but this is a
 * response/error that doesn't match a previously sent request
 * @STUN_VALIDATION_UNKNOWN_REQUEST_ATTRIBUTE: The message is valid but
 * contains one or more unknown comprehension attributes.
 * stun_agent_build_unknown_attributes_error() should be called
 * @STUN_VALIDATION_UNKNOWN_ATTRIBUTE: The message is valid but contains one
 * or more unknown comprehension attributes. This is a response, or error,
 * or indication message and no error response should be sent
 *
 * This enum is used as the return value of stun_agent_validate() and represents
 * the status result of the validation of a STUN message.
*/
type StunValidationStatus int
const (
	_ StunValidationStatus = iota
	STUN_VALIDATION_SUCCESS
	STUN_VALIDATION_NOT_STUN
	STUN_VALIDATION_INCOMPLETE_STUN
	STUN_VALIDATION_BAD_REQUEST
	STUN_VALIDATION_UNAUTHORIZED_BAD_REQUEST
	STUN_VALIDATION_UNAUTHORIZED
	STUN_VALIDATION_UNMATCHED_RESPONSE
	STUN_VALIDATION_UNKNOWN_REQUEST_ATTRIBUTE
	STUN_VALIDATION_UNKNOWN_ATTRIBUTE
)


/**
 * StunAgentUsageFlags:
 * @STUN_AGENT_USAGE_SHORT_TERM_CREDENTIALS: The agent should be using the short
 * term credentials mechanism for authenticating STUN messages
 * @STUN_AGENT_USAGE_LONG_TERM_CREDENTIALS: The agent should be using the long
 * term credentials mechanism for authenticating STUN messages
 * @STUN_AGENT_USAGE_USE_FINGERPRINT: The agent should add the FINGERPRINT
 * attribute to the STUN messages it creates.
 * @STUN_AGENT_USAGE_ADD_SOFTWARE: The agent should add the SOFTWARE attribute
 * to the STUN messages it creates. Calling nice_agent_set_software() will have
 * the same effect as enabling this Usage. STUN Indications do not have the
 * SOFTWARE attributes added to them though. The SOFTWARE attribute is only
 * added for the RFC5389 and MSICE2 compatibility modes.
 * @STUN_AGENT_USAGE_IGNORE_CREDENTIALS: The agent should ignore any credentials
 * in the STUN messages it receives (the MESSAGE-INTEGRITY attribute
 * will never be validated by stun_agent_validate())
 * @STUN_AGENT_USAGE_NO_INDICATION_AUTH: The agent should ignore credentials
 * in the STUN messages it receives if the #StunClass of the message is
 * #STUN_INDICATION (some implementation require #STUN_INDICATION messages to
 * be authenticated, while others never add a MESSAGE-INTEGRITY attribute to a
 * #STUN_INDICATION message)
 * @STUN_AGENT_USAGE_FORCE_VALIDATER: The agent should always try to validate
 * the password of a STUN message, even if it already knows what the password
 * should be (a response to a previously created request). This means that the
 * #StunMessageIntegrityValidate callback will always be called when there is
 * a MESSAGE-INTEGRITY attribute.
 * @STUN_AGENT_USAGE_NO_ALIGNED_ATTRIBUTES: The agent should not assume STUN
 * attributes are aligned on 32-bit boundaries when parsing messages and also
 * do not add padding when creating messages.
 *
 * This enum defines a bitflag usages for a #StunAgent and they will define how
 * the agent should behave, independently of the compatibility mode it uses.
 * <para> See also: stun_agent_init() </para>
 * <para> See also: stun_agent_validate() </para>
*/
type StunAgentUsageFlags int
const (
	_ StunAgentUsageFlags = iota
	STUN_AGENT_USAGE_SHORT_TERM_CREDENTIALS    = (1 << 0)
	STUN_AGENT_USAGE_LONG_TERM_CREDENTIALS     = (1 << 1)
	STUN_AGENT_USAGE_USE_FINGERPRINT           = (1 << 2)
	STUN_AGENT_USAGE_ADD_SOFTWARE              = (1 << 3)
	STUN_AGENT_USAGE_IGNORE_CREDENTIALS        = (1 << 4)
	STUN_AGENT_USAGE_NO_INDICATION_AUTH        = (1 << 5)
	STUN_AGENT_USAGE_FORCE_VALIDATER           = (1 << 6)
	STUN_AGENT_USAGE_NO_ALIGNED_ATTRIBUTES     = (1 << 7)
)

type StunAgentSavedIds struct {
	id			StunTransactionId
	StunMethod method;
	uint8_t *key;
	size_t key_len;
	uint8_t long_term_key[16];
	bool long_term_valid;
	bool valid;
}

type StunAgent struct {
	compatibility		StunCompatibility

}
