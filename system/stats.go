package system

import (
	"time"

	"github.com/perlogix/cmon/config"
	"github.com/perlogix/cmon/data"

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
