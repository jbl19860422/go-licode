package nice

type NiceSocketType int
const (
	_ NiceSocketType = iota
	NICE_SOCKET_TYPE_UDP_BSD
	NICE_SOCKET_TYPE_TCP_BSD
	NICE_SOCKET_TYPE_PSEUDOSSL
	NICE_SOCKET_TYPE_HTTP
	NICE_SOCKET_TYPE_SOCKS5
	NICE_SOCKET_TYPE_UDP_TURN
	NICE_SOCKET_TYPE_UDP_TURN_OVER_TCP
	NICE_SOCKET_TYPE_TCP_ACTIVE
	NICE_SOCKET_TYPE_TCP_PASSIVE
	NICE_SOCKET_TYPE_TCP_SO
)
type NiceSocket struct {
	addr 			NiceAddress
	typ 			NiceSocketType
	//GSocket 		*fileno
}

type NiceSocketWritableCb func(sock *NiceSockInterface, user_data interface{})

type NiceSockInterface interface {
	recv_messages(recv_msgs []*NiceInputMessage) error
	send_messages(to *NiceAddress, messages []*NiceOutputMessage) error
	send_messages_reliable(to *NiceAddress, messages []*NiceOutputMessage) error
	is_reliable() bool
	can_send(addr *NiceAddress) bool
	set_writable_callback(cb NiceSocketWritableCb)
	is_based_on(ohter *NiceSocket) bool
	close()
}


