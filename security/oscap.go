package security

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/config"
	"github.com/perlogix/cmon/data"
)

// parse reads from exec.Command StdOutPipe and converts it into JSON
func parse(s io.Reader) data.OScapOutput {
	scanner := bufio.NewScanner(s)

	processedOutput := data.OScapOutput{
		Status:   true,
		Warnings: []string{},
		Failed:   []data.OScapResult{},
	}

	res := data.OScapResult{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "WARNING") {
			processedOutput.Warnings = append(processedOutput.Warnings, line)
		}

		if strings.HasPrefix(line, "Title") {
			res.Title = strings.TrimSpace(strings.TrimPrefix(line, "Title"))
			processedOutput.Checks = processedOutput.Checks + 1
		}

		if strings.HasPrefix(line, "Rule") {
			res.Rule = strings.TrimSpace(strings.TrimPrefix(line, "Rule"))
		}

		if strings.HasPrefix(line, "Result") {
			res.Result = strings.TrimSpace(strings.TrimPrefix(line, "Result"))

			switch res.Result {
			case "pass":
				processedOutput.PassTotal++
			case "fixed":
				processedOutput.FixedTotal++
			case "informational":
				processedOutput.InfoTotal++
			case "fail":
				processedOutput.FailTotal++
			case "error":
				processedOutput.ErrorTotal++
			case "unknown":
				processedOutput.UnknownTotal++
			case "notchecked":
				processedOutput.NotCheckTotal++
			case "notselected":
				processedOutput.NotSelectTotal++
			case "notapplicable":
				processedOutput.NotAppTotal++
			}

			if isFailed(res.Result) {
				processedOutput.Status = false
				processedOutput.Failed = append(processedOutput.Failed, res)
			}

			res = data.OScapResult{}
		}
	}

	return processedOutput
}

// isFailed checks for oscap result is not pass
func isFailed(s string) bool {
	var passResults = []string{
		"pass",
		"skipped",
		"notchecked",
	}

	for _, r := range passResults {
		if s == r {
			return false
		}
	}

	return true
}

// OScap runs OpenScap Ubuntu 18.04 CIS benchmarks
func OScap(d *data.DiscoverJSON) {

	parsedOut := data.OScapOutput{
		Status:   true,
		Warnings: []string{},
		Failed:   []data.OScapResult{},
	}

	if runtime.GOOS == "linux" {

		_, err := os.Stat(config.Str("oscap_xccdf_xml"))
		if os.IsNotExist(err) {
			d.OpenScap = parsedOut
			return
		}

		oscap := exec.Command("oscap", "xccdf", "eval", "--profile", config.Str("oscap_profile"), config.Str("oscap_xccdf_xml"))

		output, err := oscap.StdoutPipe()
		if err != nil {
			d.OpenScap = parsedOut
			return
		}

		err = oscap.Start()
		if err != nil {
			d.OpenScap = parsedOut
			return
		}

		parsedOut = parse(output)
	}

	d.OpenScap = parsedOut
}
