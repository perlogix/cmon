package packages

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Deb fetches all dpkg packages
func Deb(d *data.DiscoverJSON) {
	dpkgSlice := []string{}

	if runtime.GOOS == "linux" {
		dpkgOut, _ := util.Cmd(`dpkg -l | awk '/^[a-z]/{print$2"-"$3}'`)

		dpkgSlice = append(dpkgSlice, strings.Split(strings.TrimSpace(string(dpkgOut)), "\n")...)

	}

	d.Packages = dpkgSlice
}
