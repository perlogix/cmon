package system

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// DmesgErrors detects err, emerg, crit, alert messages
func DmesgErrors(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		dmesgOut, err := util.Cmd(`dmesg -HP --level=err,emerg,crit,alert`)
		if err != nil {
			return
		}

		dmesgSlice := strings.Split(strings.TrimSpace(string(dmesgOut)), "\n")

		d.DmesgErrors = dmesgSlice
	}
}
