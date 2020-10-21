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
