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
