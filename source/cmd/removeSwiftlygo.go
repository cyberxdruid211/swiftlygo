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
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var removeSwiftlygoCmd = &cobra.Command{
	Use:   "removeSwiftlygo",
	Short: "Removes the SwiftlyGo app",
	Long: fmt.Sprintf(`%s
- removeSwiftlygo -
This command removes the SwiftlyGo app.

For example:
swiftlygo removeSwiftlygo

This command only removes the SwiftlyGo app.
All Swift installations will remain on your system.

SwiftlyGo can easily be re-installed by running the install script.
curl -sL https://swiftlygo.xyz/install.sh | sudo bash`, logo),
	Run: func(cmd *cobra.Command, args []string) {
		if os.Geteuid() != 0 {
			fmt.Printf("\n%s\n\n", adminMessage)
			os.Exit(0)
		}
		fmt.Println()
		if askForConfirmation("Are you sure you want to remove SwiftlyGo?") {
			fmt.Println("Removing swiftlygo ...")
			err := exec.Command("rm", "/usr/bin/swiftlygo").Run()
			if err != nil {
				log.Fatalf("Failed to execute command: %v", err)
			}
			fmt.Println("removed swiftlygo symlink - /usr/bin/swiftlygo")
			err = exec.Command("rm", "-rf", "/usr/libexec/swiftlygo").Run()
			fmt.Println("removed swiftlygo directory - /usr/libexec/swiftlygo")
			if err != nil {
				log.Fatalf("Failed to execute command: %v", err)
			}
			fmt.Println("SwiftlyGo removed successfully")
			fmt.Println()
		} else {
			fmt.Println("SwiftlyGo removal aborted")
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(removeSwiftlygoCmd)
}
