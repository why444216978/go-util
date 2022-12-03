package net

import "net"

func IsInternalIP(ip net.IP, ipv4Func, ipv6Func func(net.IP) bool) bool {
	if ip == nil || ipv4Func == nil || ipv6Func == nil {
		return false
	}
	if ip.IsLoopback() {
		return true
	}
	ip4 := ip.To4()
	if ip4 != nil {
		return ipv4Func(ip4)
	}
	ip6 := ip.To16()
	if ip6 != nil {
		return ipv6Func(ip6)
	}

	return false
}
