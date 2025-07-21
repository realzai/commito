package utils

import (
	"bytes"
	"os/exec"
)

func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}
