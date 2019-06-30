package nice

/**
 * NiceInputMessage:
 * @buffers: (array length=n_buffers): unowned array of #GInputVector buffers to
 * store data in for this message
 * @n_buffers: number of #GInputVectors in @buffers, or -1 to indicate @buffers
 * is %NULL-terminated
 * @from: (allow-none): return location to store the address of the peer who
 * transmitted the message, or %NULL
 * @length: total number of valid bytes contiguously stored in @buffers
 *
 * Represents a single message received off the network. For reliable
 * connections, this is essentially just an array of buffers (specifically,
 * @from can be ignored). for non-reliable connections, it represents a single
 * packet as received from the OS.
 *
 * @n_buffers may be -1 to indicate that @buffers is terminated by a
 * #GInputVector with a %NULL buffer pointer.
 *
 * By providing arrays of #NiceInputMessages to functions like
 * nice_agent_recv_messages(), multiple messages may be received with a single
 * call, which is more efficient than making multiple calls in a loop. In this
 * manner, nice_agent_recv_messages() is analogous to recvmmsg(); and
 * #NiceInputMessage to struct mmsghdr.
 *
 * Since: 0.1.5
 */
type NiceInputMessage struct {
	buffers			[]([]byte)
	from 			*NiceAddress
}

/**
 * NiceOutputMessage:
 * @buffers: (array length=n_buffers): unowned array of #GOutputVector buffers
 * which contain data to transmit for this message
 * @n_buffers: number of #GOutputVectors in @buffers, or -1 to indicate @buffers
 * is %NULL-terminated
 *
 * Represents a single message to transmit on the network. For
 * reliable connections, this is essentially just an array of
 * buffer. for non-reliable connections, it represents a single packet
 * to send to the OS.
 *
 * @n_buffers may be -1 to indicate that @buffers is terminated by a
 * #GOutputVector with a %NULL buffer pointer.
 *
 * By providing arrays of #NiceOutputMessages to functions like
 * nice_agent_send_messages_nonblocking(), multiple messages may be transmitted
 * with a single call, which is more efficient than making multiple calls in a
 * loop. In this manner, nice_agent_send_messages_nonblocking() is analogous to
 * sendmmsg(); and #NiceOutputMessage to struct mmsghdr.
 *
 * Since: 0.1.5
 */
type  NiceOutputMessage struct {
	buffers []([]byte)
}


/**
 * NICE_AGENT_MAX_REMOTE_CANDIDATES:
 *
 * A hard limit for the number of remote candidates. This
 * limit is enforced to protect against malevolent remote
 * clients.
 */
const NICE_AGENT_MAX_REMOTE_CANDIDATES = 25

/**
 * NiceComponentState:
 * @NICE_COMPONENT_STATE_DISCONNECTED: No activity scheduled
 * @NICE_COMPONENT_STATE_GATHERING: Gathering local candidates
 * @NICE_COMPONENT_STATE_CONNECTING: Establishing connectivity
 * @NICE_COMPONENT_STATE_CONNECTED: At least one working candidate pair
 * @NICE_COMPONENT_STATE_READY: ICE concluded, candidate pair selection
 * is now final
 * @NICE_COMPONENT_STATE_FAILED: Connectivity checks have been completed,
 * but connectivity was not established
 * @NICE_COMPONENT_STATE_LAST: Dummy state
 *
 * An enum representing the state of a component.
 * <para> See also: #NiceAgent::component-state-changed </para>
 */
type NiceComponentState int
const (
	_ NiceComponentState = iota
	NICE_COMPONENT_STATE_DISCONNECTED
	NICE_COMPONENT_STATE_GATHERING
	NICE_COMPONENT_STATE_CONNECTING
	NICE_COMPONENT_STATE_CONNECTED
	NICE_COMPONENT_STATE_READY
	NICE_COMPONENT_STATE_FAILED
	NICE_COMPONENT_STATE_LAST
)

/**
 * NiceComponentType:
 * @NICE_COMPONENT_TYPE_RTP: RTP Component type
 * @NICE_COMPONENT_TYPE_RTCP: RTCP Component type
 *
 * Convenience enum representing the type of a component for use as the
 * component_id for RTP/RTCP usages.
 <example>
   <title>Example of use.</title>
   <programlisting>
   nice_agent_send (agent, stream_id, NICE_COMPONENT_TYPE_RTP, len, buf);
   </programlisting>
  </example>
 */
type NiceComponentType int
const (
	_ NiceComponentType = iota
	NICE_COMPONENT_TYPE_RTP
	NICE_COMPONENT_TYPE_RTCP
)

/**
 * NiceCompatibility:
 * @NICE_COMPATIBILITY_RFC5245: Use compatibility with the RFC5245 ICE-UDP specs
 * and RFC6544 ICE-TCP specs
 * @NICE_COMPATIBILITY_GOOGLE: Use compatibility for Google Talk specs
 * @NICE_COMPATIBILITY_MSN: Use compatibility for MSN Messenger specs
 * @NICE_COMPATIBILITY_WLM2009: Use compatibility with Windows Live Messenger
 * 2009
 * @NICE_COMPATIBILITY_OC2007: Use compatibility with Microsoft Office Communicator 2007
 * @NICE_COMPATIBILITY_OC2007R2: Use compatibility with Microsoft Office Communicator 2007 R2
 * @NICE_COMPATIBILITY_DRAFT19: Use compatibility for ICE Draft 19 specs
 * @NICE_COMPATIBILITY_LAST: Dummy last compatibility mode
 *
 * An enum to specify which compatible specifications the #NiceAgent should use.
 * Use with nice_agent_new()
 *
 * <warning>@NICE_COMPATIBILITY_DRAFT19 is deprecated and should not be used
 * in newly-written code. It is kept for compatibility reasons and
 * represents the same compatibility as @NICE_COMPATIBILITY_RFC5245 </warning>
 <note>
   <para>
   If @NICE_COMPATIBILITY_RFC5245 compatibility mode is used for a non-reliable
   agent, then ICE-UDP will be used with higher priority and ICE-TCP will also
   be used when the UDP connectivity fails. If it is used with a reliable agent,
   then ICE-UDP will be used with the TCP-Over-UDP (#PseudoTcpSocket) if ICE-TCP
   fails and ICE-UDP succeeds.
  </para>
 </note>
 *
 */
type NiceCompatibility byte
const (
	_ NiceCompatibility = iota
	NICE_COMPATIBILITY_RFC5245 = 0
	NICE_COMPATIBILITY_DRAFT19 = NICE_COMPATIBILITY_RFC5245
	NICE_COMPATIBILITY_GOOGLE
	NICE_COMPATIBILITY_MSN
	NICE_COMPATIBILITY_WLM2009
	NICE_COMPATIBILITY_OC2007
	NICE_COMPATIBILITY_OC2007R2
	NICE_COMPATIBILITY_LAST = NICE_COMPATIBILITY_OC2007R2
)

/**
 * NiceProxyType:
 * @NICE_PROXY_TYPE_NONE: Do not use a proxy
 * @NICE_PROXY_TYPE_SOCKS5: Use a SOCKS5 proxy
 * @NICE_PROXY_TYPE_HTTP: Use an HTTP proxy
 * @NICE_PROXY_TYPE_LAST: Dummy last proxy type
 *
 * An enum to specify which proxy type to use for relaying.
 * Note that the proxies will only be used with TCP TURN relaying.
 * <para> See also: #NiceAgent:proxy-type </para>
 *
 * Since: 0.0.4
 */
type NiceProxyType int
const (
	_ NiceProxyType = iota
	NICE_PROXY_TYPE_NONE
	NICE_PROXY_TYPE_SOCKS5
	NICE_PROXY_TYPE_HTTP
	NICE_PROXY_TYPE_LAST = NICE_PROXY_TYPE_HTTP
)

/**
 * NiceNominationMode:
 * @NICE_NOMINATION_MODE_AGGRESSIVE: Aggressive nomination mode
 * @NICE_NOMINATION_MODE_REGULAR: Regular nomination mode
 *
 * An enum to specity the kind of nomination mode to use by
 * the agent, as described in RFC 5245. Two modes exists,
 * regular and aggressive. They differ by the way the controlling
 * agent chooses to put the USE-CANDIDATE attribute in its STUN
 * messages. The aggressive mode is supposed to nominate a pair
 * faster, than the regular mode, potentially causing the nominated
 * pair to change until the connection check completes.
 *
 * Since: 0.1.15
 */
type NiceNominationMode int
const (
	_ NiceNominationMode = iota
	NICE_NOMINATION_MODE_REGULAR
	NICE_NOMINATION_MODE_AGGRESSIVE
)

/**
 * NiceAgentOption:
 * @NICE_AGENT_OPTION_REGULAR_NOMINATION: Enables regular nomination, default
 *  is aggrssive mode (see #NiceNominationMode).
 * @NICE_AGENT_OPTION_RELIABLE: Enables reliable mode, possibly using PseudoTCP, *  see nice_agent_new_reliable().
 * @NICE_AGENT_OPTION_LITE_MODE: Enable lite mode
 * @NICE_AGENT_OPTION_ICE_TRICKLE: Enable ICE trickle mode
 * @NICE_AGENT_OPTION_SUPPORT_RENOMINATION: Enable renomination triggered by NOMINATION STUN attribute
 * proposed here: https://tools.ietf.org/html/draft-thatcher-ice-renomination-00
 *
 * These are options that can be passed to nice_agent_new_full(). They set
 * various properties on the agent. Not including them sets the property to
 * the other value.
 *
 * Since: 0.1.15
 */
type NiceAgentOption int
const (
	_ NiceAgentOption = iota
	 NICE_AGENT_OPTION_REGULAR_NOMINATION = 1 << 0
	 NICE_AGENT_OPTION_RELIABLE = 1 << 1
	 NICE_AGENT_OPTION_LITE_MODE = 1 << 2
	 NICE_AGENT_OPTION_ICE_TRICKLE = 1 << 3
	 NICE_AGENT_OPTION_SUPPORT_RENOMINATION = 1 << 4
)

/**
 * NiceAgentRecvFunc:
 * @agent: The #NiceAgent Object
 * @stream_id: The id of the stream
 * @component_id: The id of the component of the stream
 *        which received the data
 * @len: The length of the data
 * @buf: The buffer containing the data received
 * @user_data: The user data set in nice_agent_attach_recv()
 *
 * Callback function when data is received on a component
 *
*/
type NiceAgentRecvFunc func(agent *NiceAgent, stream_id uint, component_id uint,buf []byte, user_data []byte)

/**
 * nice_agent_new:
 * @ctx: The Glib Mainloop Context to use for timers
 * @compat: The compatibility mode of the agent
 *
 * Create a new #NiceAgent.
 * The returned object must be freed with g_object_unref()
 *
 * Returns: The new agent GObject
 */
func nice_agent_new(compat NiceCompatibility ) *NiceAgent {
	return &NiceAgent{
	}
}

