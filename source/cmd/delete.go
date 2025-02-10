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

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a Swift version",
	Long: fmt.Sprintf(`%s
- swiftlygo delete -
This command deletes locally installed Swift versions.

For example:

swiftlygo delete 5.10.1
This command will delete Swift version 5.10.1.

The delete command will also deactivate this version if it is currently active.`, logo),
	Run: func(cmd *cobra.Command, args []string) {
		var version string = ""
		installed_list := getInstalledVersions(installBase)
		if os.Geteuid() != 0 {
			fmt.Printf("\n%s\n\n", adminMessage)
			os.Exit(0)
		}
		if len(args) > 0 {
			version = args[0]
        } else {
			fmt.Printf("\nPlease provide a version number to %s.\n\nFor example:\n", cmd.Use)
			fmt.Printf("sudo swiftlygo %s 6.0.3\n\n", cmd.Use)
			fmt.Printf("Available versions are:\n")
			fmt.Printf("%s\n\n", strings.Join(installed_list, ", "))
			os.Exit(0)
		}

		if !checkVersionIsInstalled(version, installBase) {
			fmt.Printf("\nSorry, that version is not available.\n\nAvailable versions are:\n")
			fmt.Printf("%s\n\n", strings.Join(installed_list, ", "))
			os.Exit(0)
		}
		linkPath := linkBase+"swift"
		currentVer := getSwiftVersion(linkPath)
		if currentVer == version {
			if askForConfirmation("Version "+version+" is active! Are you sure you want to delete this version?") {
				fmt.Println("Deleting ...")
				deleteSymlinks()
				deleteDirectory(version)
				fmt.Printf("Version %s deleted.\n\n", version)
			} else {
				fmt.Printf("Delete cancelled.\n\n")
			}
		} else {
			if askForConfirmation("Are you sure you want to delete version "+version+"?") {
				fmt.Println("Deleting ...")
				deleteDirectory(version)
				fmt.Printf("Version %s deleted.\n\n", version)
			} else {
				fmt.Printf("Delete cancelled.\n\n")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
