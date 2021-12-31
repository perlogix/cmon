package network

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// IPRoutes fetches the IP routes on NIX systems
func IPRoutes(d *data.DiscoverJSON) {

	if runtime.GOOS == "linux" {
		iproute, err := exec.Command("ip", "route").Output()
		if err != nil {
			return
		}

		iprteSlice := strings.Split(strings.TrimSpace(string(iproute)), "\n")

		if iprteSlice != nil {
			d.IPRoute = iprteSlice
		}
	}

	if runtime.GOOS == "darwin" {
		net := exec.Command("netstat", "-rn")
		netGrep := exec.Command("grep", "-v", "Internet\\|Routing\\|Destination\\|^$")
		netOut, err := net.StdoutPipe()
		if err != nil {
			return
		}
		err = net.Start()
		if err != nil {
			return
		}
		netGrep.Stdin = netOut
		netsOut, err := netGrep.Output()
		if err != nil {
			return
		}

		netsSlice := strings.Split(strings.TrimSpace(string(netsOut)), "\n")

		if netsSlice != nil {
			d.IPRoute = netsSlice
		}
	}
}
