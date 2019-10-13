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

package cloud

import (
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/yeticloud/yeti-discover/data"
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
	if runtime.GOOS != "darwin" {
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
}
