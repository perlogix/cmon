package system

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Lsmod lists the currently loaded kernel modules
func Lsmod(d *data.DiscoverJSON) {
	outSlice := []string{}

	if runtime.GOOS == "linux" {

		lsmodOut, err := util.Cmd(`lsmod | grep -v Module | awk '{ print $1 }' | sort`)
		if err != nil {
			d.Lsmod = outSlice
			return
		}

		outString := strings.Split(string(lsmodOut), "\n")

		for _, s := range outString {
			if s != "" {
				space := regexp.MustCompile(`\s+`)
				s := space.ReplaceAllString(s, " ")
				outSlice = append(outSlice, s)
			}
		}
	}

	d.Lsmod = outSlice
}
