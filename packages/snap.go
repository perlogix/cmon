package packages

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Snaps fetches all snap containers
func Snaps(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		snapOut, err := util.Cmd(`snap list | awk '/^[a-z]/{print$1"-"$2}'`)
		if err != nil {
			return
		}

		snapSlice := strings.Split(strings.TrimSpace(string(snapOut)), "\n")

		d.Snaps = snapSlice
	}
}
