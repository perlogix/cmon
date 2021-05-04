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
	"net"
	"strings"
	"time"

	"github.com/perlogix/cmon/data"
)

// IP fetches the local IP address used for outbound connections
func IP(d *data.DiscoverJSON) {
	conn, err := net.DialTimeout("udp", "1.1.1.1:80", 20*time.Millisecond)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		return
	}
	d.Ipaddress = strings.Split(conn.LocalAddr().(*net.UDPAddr).String(), ":")[0]
}
