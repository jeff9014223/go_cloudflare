package tunnel

import (
	"fmt"
	"os"
)

func installTunnel(executable *[]byte) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(fmt.Sprintf("%s/.bin", home), 0755)
	if err != nil {
		return "", err
	}

	executablePath := fmt.Sprintf("%s/.bin/tunnel", home)
	err = os.WriteFile(executablePath, *executable, 0755)
	if err != nil {
		return "", err
	}

	return executablePath, nil
}
