package nice

import "net"

type NiceAddress net.Addr

func nice_address_equal_no_port (a NiceAddress, b NiceAddress) bool {
	ipnet_a, ok := a.(*net.IPNet)
	if !ok {
		return false
	}

	ipnet_b, ok2 := a.(*net.IPNet)
	if !ok2 {
		return false
	}

	return ipnet_a.IP.Equal(ipnet_b.IP)
}

func nice_address_equal (a NiceAddress, b NiceAddress) bool {
	return a.String() == b.String()
}
