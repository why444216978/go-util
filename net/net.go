package net

import (
	"context"
	"errors"
	"net"
)

func IsInternalIP(ip net.IP) bool {
	if ip == nil {
		return false
	}
	if ip.IsLoopback() {
		return true
	}
	ip4 := ip.To4()
	if ip4 != nil {
		return isInternalIP4(ip4)
	}
	ip6 := ip.To16()
	if ip6 != nil {
		return isInternalIP6(ip6)
	}

	return false
}

func isInternalIP4(ip4 net.IP) bool {
	if ip4[0] == 10 {
		return true
	}

	if ip4[0] == 192 && ip4[1] == 168 {
		return true
	}

	if ip4[0] == 172 && (ip4[1] >= 16 && ip4[1] <= 31) {
		return true
	}

	return false
}

func isInternalIP6(ip6 net.IP) bool {
	// fd00::/8
	if ip6[0] == 0xfd {
		return true
	}

	// fe80::/10
	if ip6[0] == 0xfe && (ip6[1] >= 0x80 && ip6[1] <= 0xbf) {
		return true
	}
	return false
}

func LookupIP(ctx context.Context, network, host string) ([]net.IP, error) {
	if ip := net.ParseIP(host); ip != nil {
		return []net.IP{ip}, nil
	}

	return net.DefaultResolver.LookupIP(ctx, network, host)
}

func DNSError(err error) *net.DNSError {
	var ge *net.DNSError
	errors.As(err, &ge)
	return ge
}
