package cmd

import (
	"fmt"
	"os"

	pkgcmd "github.com/kedgeproject/kedge/pkg/cmd"
	"github.com/spf13/cobra"
)

// Variables
var (
	AppFiles []string
)

// Represents the "generate" command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Kubernetes resources from App definition",
	Run: func(cmd *cobra.Command, args []string) {
		if err := pkgcmd.Generate(AppFiles); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	},
}

func init() {
	generateCmd.Flags().StringArrayVarP(&AppFiles, "files", "f", []string{}, "input files")
	RootCmd.AddCommand(generateCmd)
}
