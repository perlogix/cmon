package system

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Journalctl gets important logs
func Journalctl(d *data.DiscoverJSON) {
	outSlice := []string{}

	if runtime.GOOS == "linux" {

		journalctlOut, err := util.Cmd(`journalctl -p "emerg".."err" --no-pager -b | grep -vi 'kernel\|Logs\|ssh\|no entries'`)
		if err != nil {
			d.Journalctl = outSlice
			return
		}

		outString := strings.Split(string(journalctlOut), "\n")

		for _, s := range outString {
			if s != "" {
				space := regexp.MustCompile(`\s+`)
				s := space.ReplaceAllString(s, " ")
				outSlice = append(outSlice, s)
			}
		}
	}

	d.Journalctl = outSlice
}
