package util

import "os/exec"

func Cmd(cmd string) ([]byte, error) {

	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
