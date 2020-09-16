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

package security

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/yeticloud/yeti-discover/data"
)

// TrivyScan scans the root filesystem for vulnerabilities
func TrivyScan(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		// trivy filesystem -f json --exit-code 0 --no-progress -q /
		trivy, err := exec.Command("trivy", "filesystem", "-f", "json", "--exit-code", "0", "--no-progress", "-q", "/").Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		var data data.TrivyResults

		err = json.Unmarshal(trivy, &data)
		if err != nil {
			fmt.Print(err)
			return
		}

		d.Trivy = data
	}
}
