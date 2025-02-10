/*
Copyright © 2025 cyberxdruid211

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a Swift version",
	Long: fmt.Sprintf(`%s
- swiftlygo install -
This command installs the Swift version requested.

For example:
swiftlygo install latest
This command will install the latest Swift version.

swiftlygo install 5.10.1
This command will install Swift version 5.10.1.

If a local Swift version is already available on your system the install command can switch to and activate this version.

If you need to install the dependencies required for Swift use the following command:
swiftlygo install depends`, logo),
	Run: func(cmd *cobra.Command, args []string) {
		dist := os.Getenv("DIST_VER")
		arch := os.Getenv("ARCH_TYPE")
		versions_list := getAvailableSwiftVersionsList(dist, arch)
		var version string = ""
		if os.Geteuid() != 0 {
			fmt.Printf("\n%s\n\n", adminMessage)
			os.Exit(0)
		}
		if len(args) > 0 {
			version = args[0]
			if version == "latest"{
				version = getLatestSwiftVersion(dist, arch)
			}
        } else {
			fmt.Printf("\nPlease provide a version number to %s.\n\nFor example:\n", cmd.Use)
			fmt.Printf("sudo swiftlygo %s 6.0.3 or sudo swiftlygo %s latest\n\n", cmd.Use, cmd.Use)
			fmt.Printf("Available versions are:\n")
			fmt.Printf("%s\n\n", strings.Join(versions_list, ", "))
			os.Exit(0)
		}

		linkPath := linkBase+"swift"
		currentVer := getSwiftVersion(linkPath)
			switch version {
				case currentVer:
					fmt.Printf("Swift version %s is already installed and active.\n\n", version)
					os.Exit(0)
				case "depends":
					installSwiftDependencies()
					os.Exit(0)
				default:
					local_list := getInstalledVersions(installBase)
					if contains(local_list, version) {
						if askForConfirmation("Swift version "+version+" is available locally. Do you want to activate it?") {
							createSymlinks(version)
							fmt.Println("Swift " + version + " is now active.")
							fmt.Println()
							os.Exit(0)
						} else {
							fmt.Println("Activation cancelled.")
							fmt.Println()
							os.Exit(0)
						}
					}
				}
		if !contains(versions_list, version) {
			fmt.Printf("Sorry, that version is not available.\n\nAvailable versions are:\n")
			fmt.Printf("%s\n\n", strings.Join(versions_list, ", "))
			os.Exit(0)
		}
		fmt.Println()
		if askForConfirmation("Do you want to install Swift version "+version+"?") {
			fmt.Println("Installing ...")
			installSwift(version, installBase)
			if CheckDependencies() {
				if askForConfirmation("Some dependencies required for Swift are missing. Do you want to install them?") {
					fmt.Println()
					if err := InstallDependencies(); err != nil {
						fmt.Println(err)
					}
					fmt.Println("Swift dependency installation finished.")
					fmt.Println()
				} else {
					fmt.Printf("Dependency installation cancelled.\n\n")
				}
			}
		} else {
			fmt.Printf("Swift installation cancelled.\n\n")
			os.Exit(0)
		}
	},
}


func init() {
	rootCmd.AddCommand(installCmd)
}
