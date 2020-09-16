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
	"bufio"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/config"
	"github.com/yeticloud/yeti-discover/data"
)

// parse reads from exec.Command StdOutPipe and converts it into JSON
func parse(s io.Reader) data.OScapOutput {
	scanner := bufio.NewScanner(s)

	processedOutput := data.OScapOutput{
		Status: true,
		Failed: []data.OScapResult{},
	}

	res := data.OScapResult{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "WARNING") {
			processedOutput.Warnings = append(processedOutput.Warnings, line)
			continue
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
	if runtime.GOOS == "linux" {

		_, err := os.Stat(config.Str("oscap_xccdf_xml"))
		if os.IsNotExist(err) {
			return
		}

		oscap := exec.Command("oscap", "xccdf", "eval", "--profile", config.Str("oscap_profile"), config.Str("oscap_xccdf_xml"))

		output, err := oscap.StdoutPipe()
		if err != nil {
			return
		}

		err = oscap.Start()
		if err != nil {
			return
		}

		o := parse(output)

		d.OpenScap = o
	}
}