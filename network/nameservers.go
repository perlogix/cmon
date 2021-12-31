package network

import (
	"os/exec"
	"strings"

	"github.com/perlogix/cmon/data"
)

// DNS fetches all the DNS servers in resolv.conf
func DNS(d *data.DiscoverJSON) {

	dnsGrep := exec.Command("grep", "nameserver", "/etc/resolv.conf")
	dnsAwk := exec.Command("awk", "{print$2}")
	dnsGrepOut, err := dnsGrep.StdoutPipe()
	if err != nil {
		return
	}
	err = dnsGrep.Start()
	if err != nil {
		return
	}
	dnsAwk.Stdin = dnsGrepOut
	dnsOut, err := dnsAwk.Output()
	if err != nil {
		return
	}

	dnsSlice := strings.Split(strings.TrimSpace(string(dnsOut)), "\n")
	if dnsSlice != nil {
		d.DNSNameserver = dnsSlice
	}
}
