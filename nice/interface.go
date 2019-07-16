package nice

import (
	"net"
	"strings"
)

func nice_interfaces_get_local_ips() ([]NiceAddress, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	a := make([]NiceAddress, 0)
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
				c := NiceAddress{}
				c.family = "ip4"
				c.network = "udp"
				c.ip = strings.Split(addrs[j].String(), "/")[0]
				a = append(a, c)
			}
		}
	}
	return a, nil
}