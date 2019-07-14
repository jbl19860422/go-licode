package nice

/**
 * SECTION:stunagent
 * @short_description: STUN agent for building and validating STUN messages
 * @include: stun/stunagent.h
 * @see_also: #StunMessage
 * @stability: Stable
 *
 * The STUN Agent allows you to create and validate STUN messages easily.
 * It's main purpose is to make sure the building and validation methods used
 * are compatible with the RFC you create it with. It also tracks the transaction
 * ids of the requests you send, so you can validate if a STUN response you
 * received should be processed by that agent or not.
 *
 */


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
	id				StunTransactionId
	method			StunMethod
	key 			[]byte
	long_term_key	[8]byte
	long_term_valid	bool
	valid 			bool
}

type StunAgent struct {
	compatibility					StunCompatibility
	sent_ids						[]StunAgentSavedIds
	//uint16_t *known_attributes;
	usage_flags						StunAgentUsageFlags
	software_attribute				string
	ms_ice2_send_legacy_connchecks 	bool
}

/**
 * StunDefaultValidaterData:
 * @username: The username
 * @username_len: The length of the @username
 * @password: The password
 * @password_len: The length of the @password
 *
 * This structure is used as an element of the user_data to the
 * stun_agent_default_validater() function for authenticating a STUN
 * message during validationg.
 * <para> See also: stun_agent_default_validater() </para>
 */
type StunDefaultValidaterData struct {
	username []byte
	password []byte
}

/**
 * StunMessageIntegrityValidate:
 * @agent: The #StunAgent
 * @message: The #StunMessage being validated
 * @username: The username found in the @message
 * @username_len: The length of @username
 * @password: The password associated with that username. This argument is a
 * pointer to a byte array that must be set by the validater function.
 * @password_len: The length of @password which must also be set by the
 * validater function.
 * @user_data: Data to give the function
 *
 * This is the prototype for the @validater argument of the stun_agent_validate()
 * function.
 * <para> See also: stun_agent_validate() </para>
 * Returns: %TRUE if the authentication was successful,
 * %FALSE if the authentication failed
 */
func (this *StunAgent) StunMessageIntegrityValidate(message *StunMessage,
													username string,
													password string,
													user_data interface{}) bool {
		return true
}



