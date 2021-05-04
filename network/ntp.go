package network

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// NTPServers gets NTP servers listed in /etc/ntp.conf
func NTPServers(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
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

		var ntpSlice []string

		// timedatectl show-timesync -p ServerName --value
		// timedatectl show-timesync -p FallbackNTPServers --value

		for _, line := range strings.Split(strings.TrimSuffix(string(ntpOut), "\n"), "\n") {
			s := strings.TrimSpace(line)
			if s != "" {
				ntpSlice = append(ntpSlice, s)
			}
		}

		d.NTPServers = ntpSlice
	}

}

// NTPRunning detects if NTPD is running
func NTPRunning(d *data.DiscoverJSON) {
	if runtime.GOOS == "linx" {
		out, err := exec.Command("pidof", "ntpd").Output()
		if err != nil {
			return
		}

		// timedatectl show -p NTP --value | grep yes

		if string(out) != "" {
			d.NTPRunning = true
			return
		}
	}
}
