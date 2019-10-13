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
	"net"
	"strings"
	"time"

	"github.com/yeticloud/yeti-discover/data"
)

// IP fetches the local IP address used for outbound connections
func IP(d *data.DiscoverJSON) {
	conn, _ := net.DialTimeout("udp", "1.1.1.1:80", 10*time.Millisecond)
	if conn != nil {
		defer conn.Close()
	}
	d.Ipaddress = strings.Split(conn.LocalAddr().(*net.UDPAddr).String(), ":")[0]
}
