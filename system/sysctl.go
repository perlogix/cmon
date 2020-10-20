package system

import (
	"github.com/yeticloud/yeti-discover/data"
	"os/exec"
	"runtime"
	"strings"
)

func Sysctl(d *data.DiscoverJSON)  {
	//make sure that the operating system is not windows
	if runtime.GOOS == "windows"{
		return
	}
		cmd := exec.Command("sysctl", "-a")
		stdout, err := cmd.Output()
		if err != nil{
			return
		}
			outString := strings.Split(string(stdout), "\n")
			var outSlice []string
			var separator string
			if runtime.GOOS == "darwin" {
				separator = ":"
			}else if runtime.GOOS == "linux"{
				separator = "="
			}
			for _, s := range outString{
				if s != ""{
					key := strings.TrimSpace(strings.Split(s, separator)[0])
					value := strings.TrimSpace(strings.Split(s, separator)[1])
					outSlice = append(outSlice, key+"="+value)
				}
			}
			d.Sysctl = outSlice
}
