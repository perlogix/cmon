package system

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Audit fetches audit rules from auditctl -l Linux command
func Audit(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		auditctlOut, err := util.Cmd(`auditctl -l`)
		if err != nil {
			return
		}

		auditSlice := strings.Split(strings.TrimSpace(string(auditctlOut)), "\n")

		d.AuditRules = auditSlice
	}
}
