package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Deb fetches all dpkg packages
func Deb(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		dpkg := exec.Command("dpkg", "-l")
		dpkgAwk := exec.Command("awk", "/^[a-z]/{print$2\"-\"$3}")
		dpkgLOut, err := dpkg.StdoutPipe()
		if err != nil {
			return
		}
		err = dpkg.Start()
		if err != nil {
			return
		}
		dpkgAwk.Stdin = dpkgLOut
		dpkgOut, err := dpkgAwk.Output()
		if err != nil {
			return
		}

		dpkgSlice := strings.Split(strings.TrimSpace(string(dpkgOut)), "\n")

		if dpkgSlice != nil {
			d.Packages = dpkgSlice
		}
	}
}
