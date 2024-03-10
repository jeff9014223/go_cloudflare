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
	cmd := exec.Command(t.Executable, "tunnel", "run", "--token", t.Token)

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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	executablePath := fmt.Sprintf("%s/.bin/tunnel", homeDir)
	tunnel := Tunnel{Token: token, Stdout: stdout}

	_, err = os.Stat(executablePath)
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
		case "windows-amd64":
			url = "https://github.com/cloudflare/cloudflared/releases/download/2023.8.2/cloudflared-windows-amd64.exe"
		default:
			return nil, fmt.Errorf("unsupported os/arch GOOS: %s GOARCH: %s", goos, goarch)
		}

		bin, err := downloadTunnel(url)
		if err != nil {
			return nil, err
		}

		executable, err := installTunnel(&bin)
		if err != nil {
			return nil, err
		}

		tunnel.Executable = executable
		return &tunnel, nil
	}

	tunnel.Executable = executablePath
	return &tunnel, nil
}
