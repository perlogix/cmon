package security

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// FailedLogins runs lastb on NIX systems
func FailedLogins(d *data.DiscoverJSON) {
	lastbSlice := []string{}

	if runtime.GOOS == "windows" {
		d.FailedLogins = lastbSlice
		return
	}

	lastbOut, err := util.Cmd(`lastb | grep -v btmp`)
	if err != nil {
		d.FailedLogins = lastbSlice
		return
	}

	for _, line := range strings.Split(strings.TrimSuffix(string(lastbOut), "\n"), "\n") {
		s := strings.TrimSpace(line)
		if s != "" {
			lastbSlice = append(lastbSlice, s)
		}
	}

	d.FailedLogins = lastbSlice
}
