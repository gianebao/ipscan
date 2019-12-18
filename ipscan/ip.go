package ipscan

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// IP represents an IP
type IP struct {
	V4    string
	Hosts []string
}

// NewIPFromString creates a new IP from an IPv4 format string
func NewIPFromString(ipv4 string) IP {
	ip := IP{
		V4: ipv4,
	}

	ip.Hosts = ip.MustGetHostname()

	return ip
}

// GetSelf returns the current client's IP
func GetSelf() IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	ip := []rune(localAddr.String())

	// remove the port from the string
	return NewIPFromString(string(ip[0:strings.Index(string(ip), ":")]))
}

// MustGetHostname returns the hostnames associated to the IP and "" if error
func (ip IP) MustGetHostname() []string {
	hosts, _ := net.LookupAddr(ip.V4)
	return hosts
}

// ScanSubnet returns all IPs within the subnet
func (ip IP) ScanSubnet() []IP {
	ips := []IP{}
	parts := strings.Split(ip.V4, ".")
	subnet := fmt.Sprintf("%s.%s.%s.", parts[0], parts[1], parts[2])

	for i := 1; i < 256; i++ {
		ipv4 := fmt.Sprintf("%s%d", subnet, i)

		if ipv4 == ip.V4 {
			ips = append(ips, ip)
		} else {
			ips = append(ips, NewIPFromString(ipv4))
		}
	}

	return ips
}
