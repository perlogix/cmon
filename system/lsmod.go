package system

import (
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// TODO: lsmod | grep -v Module | awk '{ print $1 }' | sort

// Lsmod lists the currently loaded kernel modules
func Lsmod(d *data.DiscoverJSON) {
	if runtime.GOOS != "linux" {
		return
	}

	cmd, err := exec.Command("sh", "-c", "lsmod | grep -v Module | awk '{ print $1 }' | sort").Output()
	if err != nil {
		return
	}

	var (
		outSlice  []string
		outString = strings.Split(string(cmd), "\n")
	)

	for _, s := range outString {
		if s != "" {
			space := regexp.MustCompile(`\s+`)
			s := space.ReplaceAllString(s, " ")
			outSlice = append(outSlice, s)
		}
	}

	d.Lsmod = outSlice
}
