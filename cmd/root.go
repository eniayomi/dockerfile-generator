package cmd

import (
	"log"
	"os"

	"github.com/eniayomi/dockerfilegen/cmd/create"
	// "github.com/eniayomi/dockerfilegen/cmd/edit"
	// "github.com/eniayomi/dockerfilegen/cmd/get"
	// "github.com/eniayomi/dockerfilegen/cmd/run"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "gopkg.in/src-d/go-git.v4"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dockerfilegen",
		Short: "A cli tool to generate Dockerfiles",
		Long:  "A cli tool to generate Dockerfiles",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(create.CreateDockerfileCmd())
	// rootCmd.AddCommand(edit.EditCmd())
	rootCmd.AddCommand(VersionCmd())
}
