package cmd

import (
	"github.com/eniayomi/dockerfilegen/version"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	var (
		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Print current dockerfilegen version",
			Run: func(cmd *cobra.Command, args []string) {
				println(version.Version)
			},
		}
	)

	return versionCmd
}