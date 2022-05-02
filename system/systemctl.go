package system

import (
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// SystemctlFailed collects failed systemd units
func SystemctlFailed(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		systemctl := exec.Command("systemctl", "--failed", "--no-pager")
		systemctlGrep := exec.Command("grep", "-v", "^UNIT\\|listed")
		systemctlFout, err := systemctl.StdoutPipe()
		if err != nil {
			return
		}
		err = systemctl.Start()
		if err != nil {
			return
		}
		systemctlGrep.Stdin = systemctlFout
		systemctlOut, err := systemctlGrep.Output()
		if err != nil {
			return
		}

		var (
			outSlice  []string
			outString = strings.Split(string(systemctlOut), "\n")
		)
		for _, s := range outString {
			if s != "" {
				space := regexp.MustCompile(`\s+`)
				s := space.ReplaceAllString(s, " ")
				outSlice = append(outSlice, s)
			}
		}

		d.SystemctlFailed = outSlice
	}
}
