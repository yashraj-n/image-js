package utils

import "os/exec"

func NodeExists() bool {
	_, err := exec.Command("node", "-v").Output()
	if err != nil {
		return false
	}
	return true
}
