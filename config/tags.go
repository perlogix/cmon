package config

import (
	"github.com/perlogix/cmon/data"
)

// Tags parses tags in config file
func Tags(d *data.DiscoverJSON) {
	d.Tags = GetStringSlice("tags")
}
