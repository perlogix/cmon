package config

import (
	"github.com/perlogix/cmon/data"
)

// Tags parses tags in config file
func Tags(d *data.DiscoverJSON) {
	var tags []interface{}
	tags = append(tags, Str("tags"))
	d.Tags = tags
}
