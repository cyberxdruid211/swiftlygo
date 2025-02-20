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
	"path/filepath"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)
var yesFlag bool

var rootCmd = &cobra.Command{
	Use:   "swiftlygo",
	Short: "An Installer for the Swift Programming Language",
	Long: fmt.Sprintf(`%s
An Installer for the Swift Programming Language.

You can install and delete any Swift version that is available for your platform.
For example:

swiftlygo install 6.0.3
This command will install Swift version 6.0.3.

swiftlygo delete 6.0.3
This command will delete Swift version 6.0.3.

The activate command will activate a locally installed Swift version.
For example:

swiftlygo activate 6.0.3
This command will activate Swift version 6.0.3.`, logo),
CompletionOptions: cobra.CompletionOptions{
	DisableDefaultCmd: true,
},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&yesFlag, "yes", "y", false, "Automatically answer 'yes' to all prompts")
	// Get the directory of the executable
    exePath, err := os.Executable()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: could not determine executable path: %v\n", err)
        os.Exit(1) // Exit if we can't determine the path
    }
    
    // Construct the path to swiftlygo.env relative to the executable
    envFile := filepath.Join(filepath.Dir(exePath), "swiftlygo.env")
    if err := godotenv.Load(envFile); err != nil {
        fmt.Fprintf(os.Stderr, "Warning: could not load %s: %v\n", envFile, err)
    }
}
