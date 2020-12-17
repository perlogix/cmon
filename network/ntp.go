package network

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
)

// NTPServers gets NTP servers listed in /etc/ntp.conf
func NTPServers(d *data.DiscoverJSON) {

	if runtime.GOOS != "windows" {
		ntpGrep := exec.Command("grep", "^server", "/etc/ntp.conf")
		ntpAwk := exec.Command("awk", "{print$2}")
		ntpGrepOut, err := ntpGrep.StdoutPipe()
		if err != nil {
			return
		}

		err = ntpGrep.Start()
		if err != nil {
			return
		}

		ntpAwk.Stdin = ntpGrepOut
		ntpOut, err := ntpAwk.Output()
		if err != nil {
			return
		}

		ntpSlice := strings.Split(strings.TrimSpace(string(ntpOut)), "\n")
		if ntpSlice != nil {
			d.NTPServers = ntpSlice
		}
	}

}
