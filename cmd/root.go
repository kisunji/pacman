package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pacman",
	Short: "Package Manager Config Generator",
	Long: `Pacman is a CLI tool to generate configuration files from templates
for package managers in enterprise environments that have an internal repo 
or corporate proxies`,
	Version: "0.9.1-beta",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if _, err := os.Stat("template"); os.IsNotExist(err) {
		fmt.Println("template directory not found!")
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
