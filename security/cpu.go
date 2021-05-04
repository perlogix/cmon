package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// CPUvulns detects if CPUs are vulnerable
func CPUvulns(d *data.DiscoverJSON) {
	if runtime.GOOS == "linx" {
		out, err := exec.Command("grep", "-r", "Vulnerable", "/sys/devices/system/cpu/vulnerabilities").Output()
		if err != nil {
			return
		}

		d.CPUvulns = strings.Split(strings.TrimSuffix(string(out), "\n"), "\n")
	}
}
