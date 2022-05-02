package network

import (
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// DNS fetches all the DNS servers in resolv.conf
func DNS(d *data.DiscoverJSON) {

	dnsOut, err := util.Cmd(`grep nameserver /etc/resolv.conf | awk '{ print $2 }'`)
	if err != nil {
		return
	}

	dnsSlice := strings.Split(strings.TrimSpace(string(dnsOut)), "\n")
	if dnsSlice != nil {
		d.DNSNameserver = dnsSlice
	}
}
