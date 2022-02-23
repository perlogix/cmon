package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// DmesgErrors detects err, emerg, crit, alert messages
func DmesgErrors(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		out, err := exec.Command("dmesg", "-HP", "--level=err,emerg,crit,alert").Output()
		if err != nil {
			return
		}

		dmesgSlice := strings.Split(strings.TrimSpace(string(out)), "\n")
		d.DmesgErrors = dmesgSlice
	}
}
