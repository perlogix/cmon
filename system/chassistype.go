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
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/config"
	"github.com/yeticloud/yeti-discover/data"
)

// ChassisType detects the Chassis type like notebook, VM, server
func ChassisType(d *data.DiscoverJSON) {
	if runtime.GOOS != "linux" {
		return
	}
	cmd := exec.Command("dmidecode", "--string", "chassis-type")
	stdout, err := cmd.Output()
	if err != nil {
		return
	}
	inString := strings.Split(string(stdout), "\n")
	var outString string

	for _, s := range inString {
		if s != "" {
			outString = strings.ToLower(strings.TrimSpace(s))
		}
	}

	assetType := config.Str("asset_type")
	if assetType == "" {
		d.AssetType = outString
	} else {
		d.AssetType = assetType
	}

	d.ChassisType = outString
}
