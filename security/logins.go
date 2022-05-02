package security

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// FailedLogins runs lastb on NIX systems
func FailedLogins(d *data.DiscoverJSON) {
	if runtime.GOOS == "windows" {
		return
	}

	lastbOut, err := util.Cmd(`lastb | grep -v btmp`)
	if err != nil {
		return
	}

	var lastbSlice []string

	for _, line := range strings.Split(strings.TrimSuffix(string(lastbOut), "\n"), "\n") {
		s := strings.TrimSpace(line)
		if s != "" {
			lastbSlice = append(lastbSlice, s)
		}
	}

	d.FailedLogins = lastbSlice
}
