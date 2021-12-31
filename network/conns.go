package network

import (
	"fmt"
	"strconv"

	ps "github.com/mitchellh/go-ps"
	"github.com/perlogix/cmon/data"
	"github.com/shirou/gopsutil/net"
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
