package network

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// NTPServers gets NTP servers listed in /etc/ntp.conf
func NTPServers(d *data.DiscoverJSON) {
	ntpSlice := []string{}

	if runtime.GOOS == "linux" {

		ntpOut, err := util.Cmd(`grep ^server /etc/ntp.conf | awk '{ print $2 }'`)
		if err != nil {
			d.NTPServers = ntpSlice
			return
		}

		for _, line := range strings.Split(strings.TrimSuffix(string(ntpOut), "\n"), "\n") {
			s := strings.TrimSpace(line)
			if s != "" {
				ntpSlice = append(ntpSlice, s)
			}
		}

		if ntpSlice == nil {
			timectl, err := util.Cmd(`timedatectl show-timesync -p ServerName --value`)
			if err != nil {
				d.NTPServers = ntpSlice
				return
			}

			ntpSlice = strings.Split(strings.TrimSuffix(string(timectl), "\n"), "\n")
		}
	}

	d.NTPServers = ntpSlice
}

// NTPRunning detects if NTP is running
func NTPRunning(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		pidof, _ := util.Cmd(`pidof ntpd`)

		if string(pidof) != "" {
			d.NTPRunning = true
			return
		}

		timectl, _ := util.Cmd(`timedatectl show -p NTP --value`)

		if strings.Contains(strings.TrimSpace(string(timectl)), "yes") {
			d.NTPRunning = true
		}

	}
}
