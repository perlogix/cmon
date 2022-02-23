package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// CPUvulns detects if CPUs are vulnerable
func CPUVulnerabilities(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		out, err := exec.Command("grep", "-r", "Vulnerable", "/sys/devices/system/cpu/vulnerabilities").Output()
		if err != nil {
			return
		}

		d.CPUVulnerabilities = strings.Split(strings.TrimSpace(string(out)), "\n")
	}
}
