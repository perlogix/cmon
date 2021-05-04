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
