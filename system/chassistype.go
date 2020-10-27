package system

import (
	"github.com/yeticloud/yeti-discover/data"
	"os/exec"
	"runtime"
	"strings"
)

func ChassisType(d *data.DiscoverJSON) { // make sure that the operating system is linux only
	if runtime.GOOS != "linux" {
		return
	}
	cmd := exec.Command("dmidecode", "--string", "chassis-type")
	stdout, err := cmd.Output()
	if err != nil {
		return
	}
	inString := strings.Split(string(stdout), "\n")
	var outString string

	for _, s := range inString {
		if s != "" {
			outString = strings.ToLower(strings.TrimSpace(s))
		}
	}
	d.ChassisType = outString
}
