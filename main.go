package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gianebao/ipscan/ipscan"
)

func main() {

	all := flag.Bool("all", false, "lists all IPs")
	search := flag.String("search", "", "searches a matching hostname container the search word")

	flag.Parse()

	self := ipscan.GetSelf()

	print("IPV4", "Hostname")
	print("----", "--------")

	if *search != "" {
		filterHost(self, *search)
	} else if *all {
		listAll(self)
	} else {
		listAllWithHost(self)
	}
}

func filterHost(ip ipscan.IP, host string) {
	ips := ip.ScanSubnet()

	for _, ip := range ips {

		if len(ip.Hosts) == 0 {
			continue
		}

		for _, h := range ip.Hosts {
			if -1 < strings.Index(h, host) {
				print(ip.V4, strings.Join(ip.Hosts, ","))
				return
			}
		}
	}
}

func listAll(ip ipscan.IP) {
	ips := ip.ScanSubnet()

	for _, ip := range ips {
		if len(ip.Hosts) == 0 {
			print(ip.V4, "n/a")
		} else {
			print(ip.V4, strings.Join(ip.Hosts, ","))
		}
	}
}

func listAllWithHost(ip ipscan.IP) {
	ips := ip.ScanSubnet()

	for _, ip := range ips {
		if len(ip.Hosts) > 0 {
			print(ip.V4, strings.Join(ip.Hosts, ","))
		}
	}
}

func print(i ...interface{}) {
	fmt.Printf("%-30s %s\n", i...)
}
