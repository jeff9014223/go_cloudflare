package tunnel

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func DownloadTunnel(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download tunnel: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download tunnel: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read tunnel body: %w", err)
	}

	if strings.HasSuffix(url, ".tgz") {
		gzr, err := gzip.NewReader(bytes.NewReader(body))
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}

		tr := tar.NewReader(gzr)

		for {
			header, err := tr.Next()
			if err == io.EOF {
				break
			}

			if err != nil {
				return nil, fmt.Errorf("failed to read tar header: %w", err)
			}

			if header.Name == "cloudflared" {
				body, err = io.ReadAll(tr)
				if err != nil {
					return nil, fmt.Errorf("failed to read tar body: %w", err)
				}

				break
			}
		}
	}

	return body, nil
}
