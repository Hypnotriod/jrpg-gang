package util

import (
	"net"
	"net/http"

	"github.com/tomasen/realip"
)

func GetIP(r *http.Request) net.IP {
	address := realip.FromRequest(r)
	if address == "" {
		address = r.RemoteAddr
	}
	if host, _, _ := net.SplitHostPort(address); host != "" {
		address = host
	}
	ip := net.ParseIP(address)
	if ip == nil {
		return net.IPv4(0, 0, 0, 0)
	}
	return ip
}
