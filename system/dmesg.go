package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
)

// DmesgErrors detects err, emerg, crit, alert messages
func DmesgErrors(d *data.DiscoverJSON) {
	if runtime.GOOS == "linx" {
		out, err := exec.Command("dmesg", "-HP", "--level=err,emerg,crit,alert").Output()
		if err != nil {
			return
		}

		d.DmesgErrors = strings.TrimSpace(string(out))
	}
}
