package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
)

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

		for _, line := range strings.Split(strings.TrimSpace(string(lastbOutput)), "\n") {
			lastbSlice = append(lastbSlice, strings.TrimSpace(line))
		}

		d.FailedLogins = lastbSlice
	}
}
