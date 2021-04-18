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
	"os/exec"
	"runtime"

	"github.com/yeticloud/yeti-discover/data"
)

// TrivyScan scans the root filesystem for vulnerabilities
func TrivyScan(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		trivy, err := exec.Command("trivy", "-q", "filesystem", "-f", "json", "--exit-code", "0", "--no-progress", "/").Output()
		if err != nil {
			return
		}

		var trivData data.Trivy

		err = json.Unmarshal(trivy, &trivData.TrivyResults)
		if err != nil {
			return
		}

		for _, e := range trivData.TrivyResults {
			trivData.VulnToal += len(e.Vulnerabilities)
			for _, s := range e.Vulnerabilities {
				switch s.Severity {
				case "LOW":
					trivData.VulnLow++
				case "MEDIUM":
					trivData.VulnMed++
				case "HIGH":
					trivData.VulnHigh++
				case "CRITICAL":
					trivData.VulnCrit++
				case "UNKNOWN":
					trivData.VulnUnknown++
				}
			}
		}

		d.Trivy = trivData
	}
}
