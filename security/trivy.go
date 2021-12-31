package security

import (
	"encoding/json"
	"os/exec"
	"runtime"

	"github.com/perlogix/cmon/data"
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
