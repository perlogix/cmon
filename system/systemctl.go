package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
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

		d.SystemctlFailed = strings.TrimSpace(string(systemctlOut))
	}
}
