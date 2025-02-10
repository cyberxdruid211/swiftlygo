package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var dependencies = []string{
    "build-essential",
    "git",
    "gnupg2",
    "libcurl4",
    "libedit2",
    "libncurses-dev",
    "libpython3-dev",
    "libxml2",
    "libz3-dev",
    "pkg-config",
    "python3",
    "tzdata",
    "unzip",
    "zlib1g-dev",
}
// global vars
var (
    adminMessage    = "Please run swiftlygo with admin privileges. (sudo)"
    relVer = "v1.0"
)
var logo = fmt.Sprintf(`
 ____          _  __ _   _        ____       
/ ___|_      _(_)/ _| |_| |_   _ / ___| ___  
\___ \ \ /\ / / | |_| __| | | | | |  _ / _ \ 
 ___) \ V  V /| |  _| |_| | |_| | |_| | (_) |
|____/ \_/\_/ |_|_|  \__|_|\__, |\____|\___/ 
                           |___/  %s      
`, relVer)

func InstallDependencies() error {
    depsString := strings.Join(dependencies, " ")
    fmt.Printf("Installing dependencies: %s\n", depsString)
    execCmd := exec.Command("sh", "-c", "apt-get install -y "+depsString)
    execCmd.Stdout = os.Stdout
    execCmd.Stderr = os.Stderr
    if err := execCmd.Run(); err != nil {
        return fmt.Errorf("error installing dependencies: %v", err)
    }
    return nil
}
// checks if all specified packages are installed on the system.
func CheckDependencies() bool {
    for _, pkg := range dependencies {
        cmd := exec.Command("dpkg", "-s", pkg)
        var out bytes.Buffer
        cmd.Stdout = &out
        err := cmd.Run()

        if err != nil {
            if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() != 0 {
                //fmt.Printf("Package %s is not installed.\n", pkg)
                return true // Package not installed
            }
            // Handle other errors from dpkg
            //fmt.Printf("Failed to check package %s: %v\n", pkg, err)
            return true
        }

        output := out.String()
        if !strings.Contains(output, "Status: install ok installed") {
            if strings.Contains(output, "Status: deinstall ok config-files") {
                //fmt.Printf("Package %s was installed but is now uninstalled.\n", pkg)
            } else {
                //fmt.Printf("Package %s has unexpected status or not installed.\n", pkg)
            }
            return true // Package either not installed or previously installed but now uninstalled
        }
    }
    return false // All packages are installed
}
var stopSpinner = make(chan bool)

func spinner(message string) {
    for {
        select {
        case <-stopSpinner:
            fmt.Println()
            return
        default:
            for _, r := range `\|/-` {
                fmt.Printf("\r%s %c", message, r)
                time.Sleep(100 * time.Millisecond)
            }
        }
    }
}

// func updateEnvKeyValue(key string, new_value string) {
//     content, err := os.ReadFile(envBase+"swiftlygo.env")
//     if err != nil {
//         log.Fatal("Error reading .env file:", err)
//     }

//     envMap, err := godotenv.Unmarshal(string(content))
//     if err != nil {
//         log.Fatal("Error unmarshaling .env content:", err)
//     }

//     envMap[key] = new_value

//     newContent, err := godotenv.Marshal(envMap)
//     if err != nil {
//         log.Fatal("Error marshaling updated .env content:", err)
//     }

//     err = os.WriteFile(envBase+"swiftlygo.env", []byte(newContent), 0644)
//     if err != nil {
//         log.Fatal("Error writing to .env file:", err)
//     }

//     fmt.Println("Update: The latest Swift version is now "+new_value)
//     fmt.Println()
// }

func getDirNames(dirPath string) ([]string, error) {
    var dirs []string
    entries, err := os.ReadDir(dirPath)
    if err != nil {
        return nil, err
    }
    for _, entry := range entries {
        if entry.IsDir() {
            dirs = append(dirs, filepath.Base(entry.Name()))
        }
    }
    return dirs, nil
}

func getSwiftVersion(linkPath string) string{
    currentV := ""
    destination, err := os.Readlink(linkPath)
    if err != nil {
        //fmt.Printf("There is currently no installed Swift version.\n\n")
        return currentV
    }

    components := strings.Split(destination, string(filepath.Separator))
    if len(components) > 5 {
        currentV = components[len(components)-4]
    } else {
        fmt.Println("version not found")
    }
    fmt.Println()
    return currentV
}

func getAvailableSwiftVersionsList(dist, arch string) []string {
    url := "https://www.swift.org/api/v1/install/releases.json"
    response, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error fetching the list of available versions: %v", err)
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        log.Fatalf("Error: Status code %d Sorry, cannot find any information at the moment", response.StatusCode)
    }

    var swiftVersions []struct {
        Name       string `json:"name"`
        Platforms  []struct {
            Name  string   `json:"name"`
            Archs []string `json:"archs"`
        } `json:"platforms"`
    }
    if err := json.NewDecoder(response.Body).Decode(&swiftVersions); err != nil {
        log.Fatalf("Decode error: %v", err)
    }

    var availableVersions []string
    for _, version := range swiftVersions {
        for _, platform := range version.Platforms {
            if platform.Name == dist {
                for _, a := range platform.Archs {
                    if a == arch {
                        availableVersions = append(availableVersions, version.Name)
                        break
                    }
                }
            }
        }
    }

    return availableVersions
}

func getInstalledVersions(dirPath string) []string {
    dirNames, err := getDirNames(dirPath)
    if err != nil {
        return nil
    }
    return dirNames
}

func getLatestSwiftVersion(dist, arch string) string{
    url := "https://www.swift.org/api/v1/install/releases.json"
    response, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error fetching the list of available versions: %v", err)
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        log.Fatalf("Error: Status code %d Sorry, cannot find any information at the moment", response.StatusCode)
    }

    var swiftVersions SwiftVersion
    if err := json.NewDecoder(response.Body).Decode(&swiftVersions); err != nil {
        log.Fatalf("Decode error: %v", err)
    }
    latestVersion := ""
    for _, version := range swiftVersions {
        for _, platform := range version.Platforms {
            if platform.Name == dist && contains(platform.Archs, arch) {
                latestVersion = version.Name
                break
            }
        }
    }
    return latestVersion
}

func checkVersionIsInstalled(version, dirPath string) (bool) {
    dirNames, err := getDirNames(dirPath)
    if err != nil {
        return false
    }
    for _, dir := range dirNames {
        if dir == version {
            return true
        }
    }
    return false
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func installSwift(version string, installBase string) {
    installDir := version
    installPath := installBase + "/" + installDir
    filename := "/tmp/swift.tar.gz"

    if _, err := os.Stat(installPath); os.IsNotExist(err) {

        url := generateDownloadURL(version)
        log.Println("Downloading from:", url)
        if err := downloadFile(url, filename); err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("File successfully downloaded.")
        }

        err := os.MkdirAll(installPath, 0755)
        if err != nil {
            log.Fatalf("Failed to create directory: %v", err)
        }

        fmt.Println("Extracting Swift to:", installPath)
        extractSwift(installPath)

        createSymlinks(version)

        if err := os.Remove(filename); err != nil {
            fmt.Printf("failed to delete file %s: %v", filename, err)
        }
        fmt.Println("Swift " + version + " has been installed successfully.")
        fmt.Println()
    } else if err != nil {

        log.Fatalf("Unexpected error checking for directory: %v", err)
    } else {

        fmt.Println("NOTE: Swift " + version + " already exists, skipping download, just updating symlinks")

        createSymlinks(version)
        fmt.Println("Swift " + version + " is now active.")
        fmt.Println()
    }
}

func askForConfirmation(prompt string) bool {
    // If the yes flag is set, return true without asking
    if yesFlag {
        return true
    }
    reader := bufio.NewReader(os.Stdin)
    wrongCount := 0
    for {
        fmt.Printf("%s [Y/n]: ", prompt)

        response, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            return false
        }

        response = strings.ToLower(strings.TrimSpace(response))

        if response == "" || response == "y" || response == "yes" {
            return true
        } else if response == "n" || response == "no" {
            return false
        }

        wrongCount++
        if wrongCount >= 3 {
            fmt.Println("Too many invalid inputs.")
            return false
        }

        fmt.Println("Please answer with 'y' for yes or 'n' for no.")
    }
}
func installSwiftDependencies() {
    if askForConfirmation("Do you want to install the dependencies required for Swift?") {
        fmt.Println("Installing ...")
        if err := InstallDependencies(); err != nil {
            fmt.Println(err)
        }
        fmt.Println("Swift dependency installation finished.")
        fmt.Println()
    } else {
        fmt.Printf("Dependency installation cancelled.\n\n")
    }
}