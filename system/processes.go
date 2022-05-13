package system

import (
	"github.com/perlogix/cmon/data"
	"github.com/shirou/gopsutil/process"
)

// Processes collects the process table with basic metrics
func Processes(d *data.DiscoverJSON) {
	processes, err := process.Processes()
	if err != nil {
		d.Processes = []data.Processes{}
		return
	}

	procs := data.Processes{}

	for _, proc := range processes {
		pid := proc.Pid

		ppid, err := proc.Ppid()
		if err != nil {
			ppid = 00000
		}

		name, err := proc.Name()
		if err != nil {
			name = "unknown"
		}

		username, err := proc.Username()
		if err != nil {
			username = "unknown"
		}

		procs.Name = name
		procs.Pid = int(pid)
		procs.Ppid = int(ppid)
		procs.User = username
	}

	d.Processes = append(d.Processes, procs)
}
