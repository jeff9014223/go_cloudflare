package tunnel

import (
	"os"
)

func installTunnel(executable *[]byte) (string, error) {
	os.MkdirAll("/var/tmp", 0755)

	executablePath := "/var/tmp/tunnel"
	err := os.WriteFile(executablePath, *executable, 0755)
	if err != nil {
		return "", err
	}

	return executablePath, nil
}
