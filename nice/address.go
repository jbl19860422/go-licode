package nice

//type NiceAddress net.Addr

type NiceAddress struct {
	family		string
	network 	string
	ip 			string
	port 		int
}

func nice_address_equal_no_port (a NiceAddress, b NiceAddress) bool {
	//ipnet_a, ok := a.(*net.IPNet)
	//if !ok {
	//	return false
	//}
	//
	//ipnet_b, ok2 := a.(*net.IPNet)
	//if !ok2 {
	//	return false
	//}
	//
	//return ipnet_a.IP.Equal(ipnet_b.IP)
	return a.family == b.family && a.network == b.network && a.ip == b.ip
}

func nice_address_equal (a NiceAddress, b NiceAddress) bool {
	return a.family == b.family && a.network == b.network && a.ip == b.ip && a.port == b.port
}
