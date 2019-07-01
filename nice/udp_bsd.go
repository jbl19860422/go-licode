package nice

import "net"

type UdpBsdSocket struct {
	local_addr	*NiceAddress
	conn		*net.UDPConn
}

func NewUdpBsdSocket(addr *NiceAddress) *UdpBsdSocket {
	s := &UdpBsdSocket{}
	s.local_addr = addr
	return s
}

func nice_udp_bsd_socket_new(addr *NiceAddress) *UdpBsdSocket {
	return NewUdpBsdSocket(addr)
}

func (this *UdpBsdSocket) recv_messages(recv_msgs []*NiceInputMessage) error {
	return nil
}

func (this *UdpBsdSocket) send_messages(to *NiceAddress, messages []*NiceOutputMessage) error {
	return nil
}

func (this *UdpBsdSocket) send_messages_reliable(to *NiceAddress, messages []*NiceOutputMessage) error {
	return nil
}

func (this *UdpBsdSocket) is_reliable() bool {
	return false
}

func (this *UdpBsdSocket) can_send(addr *NiceAddress) bool {
	return true
}

func (this *UdpBsdSocket) set_writable_callback(cb NiceSocketWritableCb) {

}

func (this *UdpBsdSocket) is_based_on(ohter *NiceSocket) bool {
	return true
}

func (this *UdpBsdSocket) close() {

}