package nice

import "net"

func nice_interfaces_get_local_ips() ([]net.Addr, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	a := make([]net.Addr, 0)
	for i := 0; i < len(ifs); i++ {
		if ifs[i].Flags & net.FlagLoopback != 0 {
			continue;//ignore loop back
		}

		if ifs[i].Flags & net.FlagUp == 0 {
			continue;//ignore not up interface
		}
		addrs, err := ifs[i].Addrs()
		if err != nil {
			continue
		}

		for j := 0; j < len(addrs); j++ {
			ipnet, ok := addrs[j].(*net.IPNet)
			if !ok {
				continue
			}

			if ipnet.IP.To4() != nil {
				a = append(a, addrs[j])
			}
		}
	}
	return a, nil
}