package cloud

import (
	"github.com/perlogix/cmon/data"
	detectCloud "github.com/perlogix/libdetectcloud"
)

// DetectCloud detects if on cloud instance or container
func DetectCloud(d *data.DiscoverJSON) {
	d.Cloud = detectCloud.Detect()
}
