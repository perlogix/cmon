package network

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// IPRoutes fetches the IP routes on NIX systems
func IPRoutes(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		iproute, err := util.Cmd(`ip route`)
		if err != nil {
			return
		}

		iprteSlice := strings.Split(strings.TrimSpace(string(iproute)), "\n")

		if iprteSlice != nil {
			d.IPRoute = iprteSlice
		}
	}

	if runtime.GOOS == "darwin" {

		netsOut, err := util.Cmd(`netstat -rn | grep -v 'Internet\|Routing\|Destination\|^$'`)
		if err != nil {
			return
		}

		netsSlice := strings.Split(strings.TrimSpace(string(netsOut)), "\n")

		if netsSlice != nil {
			d.IPRoute = netsSlice
		}
	}
}
