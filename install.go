package tunnel

import (
	"os"
)

func installTunnel(executable *[]byte) (string, error) {
	err := os.MkdirAll("/var/tmp/tunnel", 0755)
	if err != nil {
		return "", err
	}

	executablePath := "/var/tmp/tunnel"
	err = os.WriteFile(executablePath, *executable, 0755)
	if err != nil {
		return "", err
	}

	return executablePath, nil
}
