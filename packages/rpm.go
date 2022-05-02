package packages

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Rpm fetches all RPM packages installed on the system
func Rpm(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		rpmOut, err := util.Cmd(`rpm -qa | sort`)
		if err != nil {
			return
		}

		rpmSlice := strings.Split(strings.TrimSpace(string(rpmOut)), "\n")

		d.Packages = rpmSlice
	}
}
