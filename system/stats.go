package system

import (
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

	d.CPUCount = len(cpuinfo)
	d.Memoryused = v.Used / gb
	d.Memoryfree = v.Free / gb
	d.Memorytotal = v.Total / gb
	d.Memoryusagepct = int(v.UsedPercent)
	d.Diskused = k.Used / gb
	d.Diskfree = k.Free / gb
	d.Disktotal = k.Total / gb
	d.Diskpct = int(k.UsedPercent)
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
	d.KernelArch = h.KernelArch
	d.Uptime = h.Uptime / days
	d.Environment = config.Str("environment")
	d.Public = config.Bool("public")
}
