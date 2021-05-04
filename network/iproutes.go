// Copyright (C) Perlogix
// This file is part of cmon <https://github.com/perlogix/cmon>.
//
// cmon is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cmon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with cmon.  If not, see <http://www.gnu.org/licenses/>.

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
