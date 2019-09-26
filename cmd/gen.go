package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kisunji/pacman/lib"
	"github.com/spf13/cobra"
)

var (
	user      string
	password  string
	output    string
	overwrite bool
)

// generateCmd represents the gen command
var generateCmd = &cobra.Command{
	Use:     "generate {maven|nuget|npm}",
	Aliases: []string{"gen", "g"},
	Short:   "Generate file",
	Long:    `Generates a config file from templates found in templates.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires an argument")
		}
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	ValidArgs: []string{"maven", "mvn", "nuget", "npm", "node"},
	Run:       runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&user, "user", "u", "", "Username")
	generateCmd.Flags().StringVarP(&password, "pass", "p", "", "Password")
	generateCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	generateCmd.Flags().BoolVar(&overwrite, "overwrite", false, "Overwrite existing file (if exists)")
}

func runGenerate(cmd *cobra.Command, args []string) {
	// Index 0 Guaranteed by Args validation
	switch args[0] {
	case "maven":
		handleMaven()
	case "mvn":
		handleMaven()
	default: // default should not be reached in ValidArgs is enforced
		log.Fatal("invalid argument")
	}
}

func handleMaven() {
	template := "template/settings.xml"

	var filename string
	if output == "" {
		var err error
		filename, err = lib.GetDefaultMavenConfPath()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("No output specified. File will be written to %s\n", filename)
	} else {
		filename = output
	}

	if lib.FileExists(filename) && !overwrite {
		fmt.Printf("File already exists at %s. Run with --overwrite to proceed.\n", filename)
		os.Exit(1)
	}

	bytes, err := lib.ReplaceMavenTemplate(template, user, password)
	if err != nil {
		log.Fatal(err)
	}
	err = lib.WriteMavenTemplate(bytes, filename)
	if err != nil {
		log.Fatal("error while writing", err)
	}
}
