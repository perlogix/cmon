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

		for _, line := range strings.Split(strings.TrimSuffix(string(ntpOut), "\n"), "\n") {
			s := strings.TrimSpace(line)
			if s != "" {
				ntpSlice = append(ntpSlice, s)
			}
		}

		if ntpSlice == nil {
			timectl, _ := exec.Command("timedatectl", "show-timesync", "-p", "ServerName", "--value").Output()
			ntpSlice = strings.Split(strings.TrimSuffix(string(timectl), "\n"), "\n")
		}

		d.NTPServers = ntpSlice
	}

}

// NTPRunning detects if NTP is running
func NTPRunning(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		pidof, _ := exec.Command("pidof", "ntpd").Output()

		timectl, _ := exec.Command("timedatectl", "show", "-p", "NTP", "--value").Output()

		if string(pidof) != "" {
			d.NTPRunning = true
			return
		}

		if strings.Contains(strings.TrimSpace(string(timectl)), "yes") {
			d.NTPRunning = true
		}

	}
}
