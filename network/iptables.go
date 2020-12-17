// Copyright (C) YetiCloud
// This file is part of yeti-discover <https://github.com/yeticloud/yeti-discover>.
//
// yeti-discover is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// yeti-discover is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with yeti-discover.  If not, see <http://www.gnu.org/licenses/>.

package network

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
)

// IPTables fetches all iptables rules
func IPTables(d *data.DiscoverJSON) {

	if runtime.GOOS == "linux" {
		iptableL := exec.Command("iptables", "-L")
		iptableGrep := exec.Command("grep", "-v", "^Chain\\|target\\|^$")
		iptableLOut, err := iptableL.StdoutPipe()
		if err != nil {
			return
		}
		err = iptableL.Start()
		if err != nil {
			return
		}
		iptableGrep.Stdin = iptableLOut
		iptableOut, err := iptableGrep.Output()
		if err != nil {
			return
		}

		var iptableSlice []string

		for _, line := range strings.Split(strings.TrimSuffix(string(iptableOut), "\n"), "\n") {
			iptableSlice = append(iptableSlice, strings.TrimSpace(line))
		}

		d.Iptables = iptableSlice
	}
}
