package tunnel

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Tunnel struct {
	Token      string
	Executable string
	Stdout     bool
}

func (t *Tunnel) Start() error {
	cmd := exec.Command("tunnel", "tunnel", "run", "--token", t.Token)
	cmd.Path = "/var/tmp"

	if t.Stdout {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func New(token string, stdout bool) (*Tunnel, error) {
	tunnel := Tunnel{Token: token, Stdout: stdout}

	_, err := os.Stat("/var/tmp/tunnel")
	if os.IsNotExist(err) {
		goos := runtime.GOOS
		goarch := runtime.GOARCH

		var url string
		switch fmt.Sprintf("%s-%s", goos, goarch) {
		case "darwin-amd64":
			url = "https://github.com/cloudflare/cloudflared/releases/download/2023.8.2/cloudflared-darwin-amd64.tgz"
		case "darwin-arm64":
			url = "https://github.com/cloudflare/cloudflared/releases/download/2023.8.2/cloudflared-darwin-amd64.tgz"
		case "linux-amd64":
			url = "https://github.com/cloudflare/cloudflared/releases/download/2023.8.2/cloudflared-linux-amd64"
		case "linux-arm64":
			url = "https://github.com/cloudflare/cloudflared/releases/download/2023.8.2/cloudflared-linux-arm64"
		default:
			return nil, fmt.Errorf("unsupported os/arch GOOS: %s GOARCH: %s", goos, goarch)
		}

		bin, err := downloadTunnel(url)
		if err != nil {
			return nil, err
		}

		err = installTunnel(&bin)
		if err != nil {
			return nil, err
		}

		return &tunnel, nil
	}

	return &tunnel, nil
}
