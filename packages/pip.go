package packages

import (
	"os/exec"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Pip fetches all Python 2.x packages
func Pip(d *data.DiscoverJSON) {

	pipFree := exec.Command("pip", "freeze")
	pipSort := exec.Command("sort")
	pipFreeOut, err := pipFree.StdoutPipe()
	if err != nil {
		return
	}
	err = pipFree.Start()
	if err != nil {
		return
	}
	pipSort.Stdin = pipFreeOut
	pipOut, err := pipSort.Output()
	if err != nil {
		return
	}

	pipReplace := strings.Replace(string(pipOut), "==", "-", -1)
	pipSlice := strings.Split(strings.TrimSpace(pipReplace), "\n")

	if pipSlice != nil {
		d.Pip = pipSlice
	}

}

// Pip3 fetches all Python 3.x packages
func Pip3(d *data.DiscoverJSON) {

	pipFree := exec.Command("pip3", "freeze")
	pipSort := exec.Command("sort")
	pipFreeOut, err := pipFree.StdoutPipe()
	if err != nil {
		return
	}
	err = pipFree.Start()
	if err != nil {
		return
	}
	pipSort.Stdin = pipFreeOut
	pipOut, err := pipSort.Output()
	if err != nil {
		return
	}

	pipReplace := strings.Replace(string(pipOut), "==", "-", -1)
	pipSlice := strings.Split(strings.TrimSpace(pipReplace), "\n")

	if pipSlice != nil {
		d.Pip3 = pipSlice
	}

}
