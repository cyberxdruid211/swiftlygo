package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
    installBase = "/usr/libexec/swift"
    linkBase = "/usr/bin/"
)

func generateDownloadURL(swiftVer string) string {

    urlBase := os.Getenv("URL_BASE")
    osDir := os.Getenv("OS_DIR")
    arch := os.Getenv("ARCH")
    osFileName := os.Getenv("OS_FILE_NAME")

    return fmt.Sprintf("%s/swift-%s-release/%s%s/swift-%s-RELEASE/swift-%s-RELEASE-%s%s.tar.gz", urlBase, swiftVer, osDir, arch, swiftVer, swiftVer, osFileName, arch)
}

func downloadFile(url, filename string) error {
    go spinner("Downloading Swift ...")
    out, err := os.Create(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to create file %s: %v\n", filename, err)
        os.Exit(1)
    }
    defer out.Close()

    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to download from %s: %v\n", url, err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Fprintf(os.Stderr, "sorry version not available: %s\n", resp.Status)
        os.Exit(1)
    }

    _, err = io.Copy(out, resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to save response body to file %s: %v\n", filename, err)
        os.Exit(1)
    }
    stopSpinner <- true
    return nil
}

func extractSwift(installDir string) {
    go spinner("Extracting Swift file ...")
    err := exec.Command("tar", "-xzf", "/tmp/swift.tar.gz", "--strip-components=1", "-C", installDir).Run()
		if err != nil {
			log.Fatalf("Failed to execute command: %v", err)
		}
    stopSpinner <- true
}

func createSymlinks(version string) {

    err := exec.Command("ln", "-sf", installBase+"/"+version+"/usr/bin/swift", linkBase+"swift").Run()
    if err != nil {
        log.Fatalf("Failed to create symlink: %v", err)
    }
    err = exec.Command("ln", "-sf", installBase+"/"+version+"/usr/bin/swiftc", linkBase+"swiftc").Run()
    if err != nil {
        log.Fatalf("Failed to create symlink: %v", err)
    }
    err = exec.Command("ln", "-sf", installBase+"/"+version+"/usr/bin/sourcekit-lsp", linkBase+"sourcekit-lsp").Run()
    if err != nil {
        log.Fatalf("Failed to create symlink: %v", err)
    }
}

func deleteDirectory(version string) error {
    dirPath := filepath.Join(installBase, version)
    err := os.RemoveAll(dirPath)
    if err != nil {
        return fmt.Errorf("failed to remove directory %s: %w", dirPath, err)
    }
    return nil
}

func deleteSymlinks() {
    err := exec.Command("rm", linkBase+"swift").Run()
    if err != nil {
        log.Fatalf("Failed to delete symlink: %v", err)
    }
    err = exec.Command("rm", linkBase+"swiftc").Run()
    if err != nil {
        log.Fatalf("Failed to delete symlink: %v", err)
    }
    err = exec.Command("rm", linkBase+"sourcekit-lsp").Run()
    if err != nil {
        log.Fatalf("Failed to delete symlink: %v", err)
    }
}
