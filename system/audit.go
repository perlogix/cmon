package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Audit fetches audit rules from auditctl -l Linux command
func Audit(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		audit, err := exec.Command("auditctl", "-l").Output()
		if err != nil {
			return
		}

		auditSlice := strings.Split(strings.TrimSpace(string(audit)), "\n")

		if auditSlice != nil {
			d.AuditRules = auditSlice
		}
	}
}
