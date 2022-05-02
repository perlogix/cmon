package packages

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Deb fetches all dpkg packages
func Deb(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		dpkgOut, err := util.Cmd(`dpkg -l | awk '/^[a-z]/{print$2"-"$3}'`)
		if err != nil {
			return
		}

		dpkgSlice := strings.Split(strings.TrimSpace(string(dpkgOut)), "\n")

		if dpkgSlice != nil {
			d.Packages = dpkgSlice
		}
	}
}
