package system

import (
	"github.com/shirou/gopsutil/process"
	"github.com/yeticloud/yeti-discover/data"
	"strconv"
	"strings"
)

func Processes(d *data.DiscoverJSON) {
	processes, err := process.Processes()

	if err != nil {
		return
	}
	var (
		outSlice []string
	)
	for _, proc := range processes {
		pid := proc.Pid

		ppid, _ := proc.Ppid()

		name, err := proc.Name()
		if err != nil {
			name = "UNKNOWN"
		}

		username, err := proc.Username()
		if err != nil {
			username = "UNKNOWN"
		}

		var cpuPct int
		cpuFloat, err := proc.CPUPercent()
		if err == nil {
			cpuPct = int(cpuFloat)
		}

		var memPct int
		memFloat, err := proc.MemoryPercent()
		if err == nil {
			memPct = int(memFloat)
		}
		s := " pid=" + strconv.Itoa(int(pid)) + " ppid=" + strconv.Itoa(int(ppid)) + " name=" + name + " user=" + username + " cpu_pct=" + strconv.Itoa(cpuPct) + " mem_pct=" + strconv.Itoa(memPct)
		s = strings.TrimSpace(s)
		outSlice = append(outSlice, s)
	}
	d.Processes = outSlice
}
