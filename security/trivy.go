package security

import (
	"encoding/json"
	"os/exec"
	"runtime"

	"github.com/aquasecurity/trivy/pkg/types"
	"github.com/perlogix/cmon/data"
)

// TrivyScan scans the root filesystem for vulnerabilities
func TrivyScan(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		trivy, _ := exec.Command("trivy", "-q", "fs", "-f", "json", "--offline-scan", "--no-progress", "--skip-policy-update", "--skip-update", "--security-checks", "vuln", "/").Output()

		var trivData data.Trivy
		var trivReport types.Report

		err := json.Unmarshal(trivy, &trivReport)
		if err != nil {
			return
		}

		for _, e := range trivReport.Results {
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

		trivData.TrivyResults = trivReport.Results

		d.Trivy = trivData
	}
}
