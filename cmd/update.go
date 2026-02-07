package cmd

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Orion to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Checking for updates...")

		release, err := getLatestRelease()
		if err != nil {
			return fmt.Errorf("failed to fetch release info: %w", err)
		}

		if release.TagName == Version {
			fmt.Printf("Orion is up to date (%s)\n", Version)
			return nil
		}

		fmt.Printf("New version available: %s\n", release.TagName)
		fmt.Print("Updating...")

		assetUrl, err := findAssetURL(release.Assets)
		if err != nil {
			return fmt.Errorf("failed to find compatible asset: %w", err)
		}

		if err := doUpdate(assetUrl); err != nil {
			return fmt.Errorf("update failed: %w", err)
		}

		fmt.Printf("\nSuccessfully updated to %s!\n", release.TagName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func getLatestRelease() (*Release, error) {
	resp, err := http.Get("https://api.github.com/repos/TanmayDabhade/orion/releases/latest")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned %s", resp.Status)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}
	return &release, nil
}

func findAssetURL(assets []Asset) (string, error) {
	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Map Go arch to GoReleaser naming convention
	// .goreleaser.yaml: 
	//   - if eq .Arch "amd64" }}x86_64
	//   - else if eq .Arch "386" }}i386
	//   - else }}{{ .Arch }}{{ end }}
	
	var targetArch string
	switch arch {
	case "amd64":
		targetArch = "x86_64"
	case "386":
		targetArch = "i386"
	default:
		targetArch = arch
	}

	// Prepare search terms
	// Title case for OS: darwin -> Darwin, linux -> Linux
	targetOS := strings.Title(osName)

	// Expected filename format: orion_Darwin_x86_64.tar.gz
	// We look for partial matches to be safe
	for _, asset := range assets {
		if strings.Contains(asset.Name, targetOS) && strings.Contains(asset.Name, targetArch) && strings.HasSuffix(asset.Name, ".tar.gz") {
			return asset.BrowserDownloadURL, nil
		}
	}

	return "", fmt.Errorf("no asset found for %s/%s", osName, arch)
}

func doUpdate(url string) error {
	// 1. Download tarball
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 2. Extract binary from tarball
	gzr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	
	var binaryData []byte
	found := false

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Look for the binary named "orion"
		if header.Typeflag == tar.TypeReg && header.Name == "orion" {
			binaryData, err = io.ReadAll(tr)
			if err != nil {
				return err
			}
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("binary 'orion' not found in update archive")
	}

	// 3. Replace current binary
	currentExe, err := os.Executable()
	if err != nil {
		return err
	}
	
	// Resolve symlinks if any (like Homebrew symlinks)
	realExe, err := filepath.EvalSymlinks(currentExe)
	if err != nil {
		// Callback to original path if eval fails
		realExe = currentExe
	}

	// Create temp file
	tmpFile, err := os.CreateTemp(filepath.Dir(realExe), ".orion-new-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name()) // Clean up if something fails

	if _, err := tmpFile.Write(binaryData); err != nil {
		return err
	}
	if err := tmpFile.Chmod(0755); err != nil {
		return err
	}
	tmpFile.Close()

	// Atomic replace
	if err := os.Rename(tmpFile.Name(), realExe); err != nil {
		return err
	}

	return nil
}
