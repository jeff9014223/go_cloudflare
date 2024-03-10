package tunnel

import (
	"os"
)

func installTunnel(executable *[]byte) error {
	os.MkdirAll("/var/tmp", 0755)

	err := os.WriteFile("/var/tmp/tunnel", *executable, 0755)
	if err != nil {
		return err
	}

	return nil
}
