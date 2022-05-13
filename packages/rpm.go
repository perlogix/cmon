package packages

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Rpm fetches all RPM packages installed on the system
func Rpm(d *data.DiscoverJSON) {
	rpmSlice := []string{}

	if runtime.GOOS == "linux" {

		rpmOut, _ := util.Cmd(`rpm -qa | sort`)

		rpmSlice = append(rpmSlice, strings.Split(strings.TrimSpace(string(rpmOut)), "\n")...)
	}

	d.Packages = rpmSlice
}
