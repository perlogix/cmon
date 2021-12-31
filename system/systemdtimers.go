package system

import (
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// SystemdTimers captures all cron like jobs within Systemd
func SystemdTimers(d *data.DiscoverJSON) {
	if runtime.GOOS != "linux" {
		return
	}
	c1 := exec.Command("systemctl", "list-timers", "--all", "--no-pager")
	c2 := exec.Command("grep", "-v", "NEXT\\|listed")
	stdout1, err := c1.StdoutPipe()
	if nil != err {
		return
	}
	err = c1.Start()
	if nil != err {
		return
	}
	c2.Stdin = stdout1
	stdout, err := c2.Output()
	if err != nil {
		return
	}

	var (
		outSlice  []string
		outString = strings.Split(string(stdout), "\n")
	)

	for _, s := range outString {
		if s != "" {
			space := regexp.MustCompile(`\s+`)
			s := space.ReplaceAllString(s, " ")
			s = strings.TrimSpace(s)
			outSlice = append(outSlice, s)
		}
	}
	d.SystemdTimers = outSlice
}
