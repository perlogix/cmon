package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Snaps fetches all snap containers
func Snaps(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		snap := exec.Command("snap", "list")
		snapAwk := exec.Command("awk", "/^[a-z]/{print$1\"-\"$2}")
		snapOut, err := snap.StdoutPipe()
		if err != nil {
			return
		}
		err = snap.Start()
		if err != nil {
			return
		}
		snapAwk.Stdin = snapOut
		snapLOut, err := snapAwk.Output()
		if err != nil {
			return
		}

		snapSlice := strings.Split(strings.TrimSpace(string(snapLOut)), "\n")

		if snapSlice != nil {
			d.Snaps = snapSlice
		}
	}
}
