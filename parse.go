package iptree

import (
	"fmt"
	"net"
	"strings"
)

type IPRange = [2]net.IP

func Parse(s string) (*net.IP, *IPRange, error) {
	// ip range type, like: "1.1.1.1-2.2.2.2"
	if strings.Contains(s, "-") {
		ips := strings.Split(s, "-")
		if len(ips) != 2 {
			return nil, nil, fmt.Errorf(
				"ip range need 2 IP, and seperate with '-'")
		}
		a := net.ParseIP(ips[0])
		if a == nil {
			return nil, nil, fmt.Errorf("head ip parse failed: %s", s)
		}
		b := net.ParseIP(ips[1])
		if b == nil {
			return nil, nil, fmt.Errorf("tail ip parse failed: %s", s)
		}
		return nil, &[2]net.IP{a, b}, nil
	}

	// ip net
	if strings.Contains(s, "/") {
		_, ipnet, err := net.ParseCIDR(s)
		if err != nil || ipnet == nil {
			return nil, nil, fmt.Errorf("%s parse failed with CIDR format", s)
		}

		a := ipnet.IP
		b := make(net.IP, len(a))
		copy(b, a)
		for i, x := range a {
			b[i] = ^(x ^ ipnet.Mask[i])
		}

		return nil, &[2]net.IP{a.To16(), b.To16()}, nil
	}

	// single ip
	a := net.ParseIP(s)
	if a == nil {
		return nil, nil, fmt.Errorf("%s parse failed with IP format", s)
	}
	return &a, nil, nil
}
