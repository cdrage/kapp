package cmd

import (
	"fmt"
	"os"

	pkgcmd "github.com/kedgeproject/kedge/pkg/cmd"
	"github.com/spf13/cobra"
)

// Variables
var (
	DeployFiles []string
)

// Represents the "deploy" command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an application to Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if err := pkgcmd.Deploy(DeployFiles); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	},
}

func init() {
	deployCmd.Flags().StringArrayVarP(&DeployFiles, "files", "f", []string{}, "Specify files")
	RootCmd.AddCommand(deployCmd)
}
