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
	"strconv"
	"strings"

	"github.com/perlogix/cmon/data"
)

// IfaceStats fetches the Kernel Network Interface Table
func IfaceStats(d *data.DiscoverJSON) {
	if runtime.GOOS != "windows" {
		net := exec.Command("netstat", "-i")
		netGrep := exec.Command("grep", "-v", "Iface\\|Kern")
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

		var ifSlice = []data.IfaceData{}

		for _, line := range strings.Split(strings.TrimSuffix(string(netsOut), "\n"), "\n") {

			fields := strings.Fields(line)
			mtu, _ := strconv.Atoi(fields[1])
			rxok, _ := strconv.Atoi(fields[2])
			rxerr, _ := strconv.Atoi(fields[3])
			rxdrp, _ := strconv.Atoi(fields[4])
			rxovr, _ := strconv.Atoi(fields[5])
			txok, _ := strconv.Atoi(fields[6])
			txerr, _ := strconv.Atoi(fields[7])
			txdrp, _ := strconv.Atoi(fields[8])
			txovr, _ := strconv.Atoi(fields[9])

			ifSlice = append(ifSlice, data.IfaceData{
				Interface: fields[0],
				MTU:       mtu,
				RXok:      rxok,
				RXerr:     rxerr,
				RXdrp:     rxdrp,
				RXovr:     rxovr,
				TXok:      txok,
				TXerr:     txerr,
				TXdrp:     txdrp,
				TXovr:     txovr,
				Flag:      fields[10],
			})
		}
		d.Interfaces = ifSlice
	}
}
