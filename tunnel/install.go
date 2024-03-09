package tunnel

import (
	"fmt"
	"os"
)

func installTunnel(executable *[]byte) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	err = os.MkdirAll(fmt.Sprintf("%s/.cloudflared", home), 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create .cloudflared directory: %w", err)
	}

	executablePath := fmt.Sprintf("%s/.cloudflared/tunnel", home)
	err = os.WriteFile(executablePath, *executable, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to write tunnel executable: %w", err)
	}

	return executablePath, nil
}
