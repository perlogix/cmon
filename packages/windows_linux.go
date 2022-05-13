package packages

import "github.com/perlogix/cmon/data"

func WindowsPackages(d *data.DiscoverJSON) {
	wp := []data.WindowsPackages{}
	d.WindowsPackages = wp
}
