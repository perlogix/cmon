package cloud

import (
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/perlogix/cmon/data"
)

var c = &http.Client{Timeout: 10 * time.Millisecond}

func awsClient(route string) string {
	url := "http://169.254.169.254/latest/" + route
	resp, err := c.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

// AWS grabs meta-data from AWS instance
func AWS(d *data.DiscoverJSON) {
	if runtime.GOOS == "darwin" {
		return
	}

	awsResponse, err := c.Get("http://169.254.169.254/latest/")
	if err != nil {
		return
	}

	if awsResponse != nil && awsResponse.StatusCode == 200 {
		d.Ec2AmiID = awsClient("ami-id")
		d.Ec2InstanceID = awsClient("instance-id")
		d.Ec2InstanceType = awsClient("instance-type")
		d.Ec2AvailabilityZone = awsClient("placement/availability-zone")
		d.Ec2Profile = awsClient("profile")
		d.Ec2PublicIP4 = awsClient("public-ipv4")
		d.Ec2SecurityGroups = strings.Split(strings.TrimSpace(awsClient("security-groups")), "\n")
	}
}
