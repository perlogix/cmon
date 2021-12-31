package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Rpm fetches all RPM packages installed on the system
func Rpm(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		rpmqa := exec.Command("rpm", "-qa")
		rpmSort := exec.Command("sort")
		rpmqaOut, err := rpmqa.StdoutPipe()
		if err != nil {
			return
		}
		err = rpmqa.Start()
		if err != nil {
			return
		}
		rpmSort.Stdin = rpmqaOut
		rpmOut, err := rpmSort.Output()
		if err != nil {
			return
		}

		rpmSlice := strings.Split(strings.TrimSpace(string(rpmOut)), "\n")

		if rpmSlice != nil {
			d.Packages = rpmSlice
		}
	}
}
