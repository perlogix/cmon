package system

import (
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Lsmod lists the currently loaded kernel modules
func Lsmod(d *data.DiscoverJSON) {
	if runtime.GOOS != "linux" {
		return
	}
	c1 := exec.Command("lsmod")
	c2 := exec.Command("grep", "-v", "Module")
	stdout1, err := c1.StdoutPipe()
	if nil != err {
		return
	}
	err = c1.Start()
	if nil != err {
		return
	}
	c2.Stdin = stdout1
	stdout, err := c2.Output()
	if err != nil {
		return
	}
	var (
		outSlice  []string
		outString = strings.Split(string(stdout), "\n")
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
