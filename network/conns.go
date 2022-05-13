package network

import (
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

	p := []data.OpenPorts{}

	c, err := net.Connections("inet")
	if err != nil {
		d.OpenPorts = p
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
			ports := data.OpenPorts{
				Address:  i.Laddr.IP,
				Port:     int(i.Laddr.Port),
				Name:     psName,
				Protocol: inetType[i.Type],
			}
			p = append(p, ports)
		}
	}

	d.OpenPorts = p
}
