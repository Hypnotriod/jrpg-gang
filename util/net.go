package util

import (
	"net"
)

func GetIP(address string) net.IP {
	host, _, _ := net.SplitHostPort(address)
	if host != "" {
		address = host
	}
	ip := net.ParseIP(address)
	if ip == nil {
		return net.IPv4(0, 0, 0, 0)
	}
	return ip
}
