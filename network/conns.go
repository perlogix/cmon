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
	"fmt"
	"strconv"

	ps "github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/net"
	"github.com/yeticloud/yeti-discover/data"
)

// Conns fetches ports open with listening address and other process information
func Conns(d *data.DiscoverJSON) {
	inetType := map[uint32]string{
		1: "tcp",
		2: "udp",
	}
	var p []string
	c, err := net.Connections("inet")
	if err != nil {
		return
	}
	for _, i := range c {
		if i.Status == "LISTEN" {
			var psName string
			proc, err := ps.FindProcess(int(i.Pid))
			if err != nil || proc == nil {
				psName = strconv.FormatInt(int64(i.Pid), 10)
			} else {
				psName = proc.Executable()
			}
			p = append(p, fmt.Sprintf("addr=%s port=%d name=%s proto=%s", i.Laddr.IP, i.Laddr.Port, psName, inetType[i.Type]))
		}
	}
	d.OpenPorts = p
}
