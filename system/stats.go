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
	"time"

	"github.com/yeticloud/yeti-discover/config"
	"github.com/yeticloud/yeti-discover/data"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	gb   = 1073741824
	days = 86400
)

func isVirt(v string) bool {
	return v == "host"
}

// Stats discovers host metrics
func Stats(d *data.DiscoverJSON) {
	cpuinfo, _ := cpu.Info()
	v, _ := mem.VirtualMemory()
	k, _ := disk.Usage("/")
	h, _ := host.Info()
	l, _ := load.Avg()
	memUsed := v.Used / gb
	memFree := v.Free / gb
	memTotal := v.Total / gb
	diskUsed := k.Used / gb
	diskFree := k.Free / gb
	diskTotal := k.Total / gb

	var cpuPct int
	cpuArr, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		cpuPct = 0
	} else {
		var total float64
		for _, value := range cpuArr {
			total += value
		}
		cpuPct = int(total)
	}

	d.CPUCount = len(cpuinfo)
	d.CPUPct = cpuPct
	d.Memoryused = memUsed
	d.Memoryfree = memFree
	d.Memorytotal = memTotal
	d.Diskused = diskUsed
	d.Diskfree = diskFree
	d.Disktotal = diskTotal
	d.Load1 = l.Load1
	d.Load5 = l.Load5
	d.Load15 = l.Load15
	d.Hostname = h.Hostname
	d.Platform = h.Platform
	d.Platformfamily = h.PlatformFamily
	d.Platformversion = h.PlatformVersion
	d.Kernelversion = h.KernelVersion
	d.Virtualization = isVirt(h.VirtualizationRole)
	d.Virtualizationsystem = h.VirtualizationSystem
	d.Os = h.OS
	d.Uptime = h.Uptime / days
	d.Environment = config.Str("environment")
	d.Public = config.Bool("public")
}
