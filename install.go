package tunnel

import (
	"os"
)

func installTunnel(executable *[]byte) error {
	err := os.WriteFile("/var/tmp/tunnel", *executable, 0755)
	if err != nil {
		return err
	}

	return nil
}
