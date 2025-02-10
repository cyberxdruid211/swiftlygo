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

type SwiftVersion []struct {
	Name           string `json:"name"`
	Tag            string `json:"tag"`
	Date           string `json:"date"`
	Platforms      []struct {
		Name     string   `json:"name"`
		Platform string   `json:"platform"`
		Archs    []string `json:"archs"`
	} `json:"platforms"`
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the Swift versions available",
	Long: fmt.Sprintf(`%s
- swiftlygo list -
This command lists the Swift versions available that can be installed on your platform.
It will also show the Swift versions that are already downloaded on your system
and the version currently active.
For example:

Swift versions available for installation on this system:-
5.10.1, 6.0, 6.0.1, 6.0.2, 6.0.3

Locally available versions:-
5.10.1, 6.0.3

The Swift version currently active is - 6.0.3
`, logo),
	Run: func(cmd *cobra.Command, args []string) {
		dist := os.Getenv("DIST_VER")
		arch := os.Getenv("ARCH_TYPE")

		if os.Geteuid() != 0 {
			fmt.Printf("\n%s\n\n", adminMessage)
			os.Exit(0)
		}

		fmt.Println()
		fmt.Printf("OS System is compatible with Swift for %s / %s \n\n", dist, arch)

		versions_list := getAvailableSwiftVersionsList(dist, arch)
		fmt.Println("Swift versions available for installation are:-")
		fmt.Print(strings.Join(versions_list, ", "))
		fmt.Printf("\n\n")

		installed_list := getInstalledVersions(installBase)
		if len(installed_list) != 0 {
			fmt.Println("Swift versions available locally are:-")
			fmt.Print(strings.Join(installed_list, ", "))
			fmt.Println()
			linkPath := linkBase+"swift"
			currentVer := getSwiftVersion(linkPath)
			if currentVer != "" {
				fmt.Printf("The Swift version currently active is - %s\n\n", currentVer)
			} else {
				fmt.Println("There is no Swift version currently active.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
