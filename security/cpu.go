package security

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// CPUvulns detects if CPUs are vulnerable
func CPUVulnerabilities(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		cpuVulns, err := util.Cmd(`grep -r Vulnerable /sys/devices/system/cpu/vulnerabilities`)
		if err != nil {
			return
		}

		d.CPUVulnerabilities = strings.Split(strings.TrimSpace(string(cpuVulns)), "\n")
	}
}
