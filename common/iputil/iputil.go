package iputil

import "net"

// IsCidr determines if the given ip is a cidr range
func IsCidr(ip string) bool {
	_, _, err := net.ParseCIDR(ip)
	if err != nil {
		return false
	}

	return true
}

// IsIP determines if the given string is a valid ip
func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
