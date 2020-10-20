package system

import (
	"github.com/yeticloud/yeti-discover/data"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func Lsmod(d *data.DiscoverJSON) { // make sure that the operating system is linux only
	if runtime.GOOS != "linux" {
		return
	}
	var cmd = exec.Command("bash", "-c", "lsmod | grep -v Module")
	stdout, err := cmd.Output()
	if nil != err {
		return
	}
	var outString = strings.Split(string(stdout), "\n")
	var (
		outSlice []string
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
