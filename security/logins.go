package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// FailedLogins runs lastb on NIX systems
func FailedLogins(d *data.DiscoverJSON) {
	if runtime.GOOS != "windows" {
		lastb := exec.Command("lastb")
		lastbGrep := exec.Command("grep", "-v", "btmp")
		lastbOut, err := lastb.StdoutPipe()
		if err != nil {
			return
		}
		err = lastb.Start()
		if err != nil {
			return
		}
		lastbGrep.Stdin = lastbOut
		lastbOutput, err := lastbGrep.Output()
		if err != nil {
			return
		}

		var lastbSlice []string

		for _, line := range strings.Split(strings.TrimSuffix(string(lastbOutput), "\n"), "\n") {
			s := strings.TrimSpace(line)
			if s != "" {
				lastbSlice = append(lastbSlice, s)
			}
		}

		d.FailedLogins = lastbSlice
	}
}
