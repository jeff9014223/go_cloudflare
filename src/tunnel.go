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

func (t *Tunnel) Start() {
	cmd := exec.Command(t.Executable, "tunnel", "run", "--token", t.Token)

	if t.Stdout {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	fmt.Println("[ Info ] Started Cloudflare Tunnel")
	cmd.Run()
}

func New(token string, stdout bool) *Tunnel {
	homeDir, _ := os.UserHomeDir()
	executablePath := fmt.Sprintf("%s/.cloudflared/tunnel", homeDir)
	tunnel := Tunnel{Token: token, Stdout: stdout}

	_, err := os.Stat(executablePath)
	if os.IsNotExist(err) {
		fmt.Println("[ Debug ] Tunnel not found, installing...")
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
			fmt.Println("Unsupported OS/Architecture:", goos, goarch)
			return nil
		}

		bin, _ := DownloadTunnel(url)
		executable, _ := InstallTunnel(&bin)
		tunnel.Executable = executable

		return &tunnel
	}

	tunnel.Executable = executablePath
	return &tunnel
}
