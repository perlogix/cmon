// Copyright (C) YetiCloud
// This file is part of yeti-discover <https://github.com/yeticloud/yeti-discover>.
//
// yeti-discover is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// yeti-discover is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with yeti-discover.  If not, see <http://www.gnu.org/licenses/>.

package system

import (
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
	"github.com/yeticloud/yeti-discover/data"
)

// Processes collects the process table with basic metrics
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
