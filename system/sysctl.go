package system

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Sysctl collects system sysctl kernel parameters
func Sysctl(d *data.DiscoverJSON) {

	outSlice := []string{}

	if runtime.GOOS == "windows" {
		d.Sysctl = outSlice
		return
	}

	sysctlOut, err := util.Cmd(`sysctl -a`)
	if err != nil {
		d.Sysctl = outSlice
		return
	}

	outString := strings.Split(string(sysctlOut), "\n")

	var separator string

	if runtime.GOOS == "darwin" {
		separator = ":"
	} else if runtime.GOOS == "linux" {
		separator = "="
	}

	for _, s := range outString {
		if s != "" {
			key := strings.TrimSpace(strings.Split(s, separator)[0])
			value := strings.TrimSpace(strings.Split(s, separator)[1])
			outSlice = append(outSlice, key+"="+value)
		}
	}

	d.Sysctl = outSlice
}
