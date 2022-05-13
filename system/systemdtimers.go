package system

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// SystemdTimers captures all cron like jobs within Systemd
func SystemdTimers(d *data.DiscoverJSON) {
	outSlice := []string{}

	if runtime.GOOS == "linux" {

		systemctlOut, err := util.Cmd(`systemctl list-timers --all --no-pager | grep -v 'NEXT\|listed'`)
		if err != nil {
			d.SystemdTimers = outSlice
			return
		}

		outString := strings.Split(string(systemctlOut), "\n")

		for _, s := range outString {
			if s != "" {
				space := regexp.MustCompile(`\s+`)
				s := space.ReplaceAllString(s, " ")
				s = strings.TrimSpace(s)
				outSlice = append(outSlice, s)
			}
		}
	}

	d.SystemdTimers = outSlice
}
