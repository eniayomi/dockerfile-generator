package create

import (
	"github.com/spf13/cobra"
)

func CreateDockerfileCmd() *cobra.Command {

	var (
		languageFlag    string

		createDockerfileCmd = &cobra.Command{
			Use:   "create",
			Short: "Create a new Dockerfile",
			Long:  "Create a new Dockerfile",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				resource.CreateProject(args[0], languageFlag, descriptionFlag, createRepoFlag)
			},
		}
	)
	createProjectCmd.Flags().StringVarP(&languageFlag, "language", "l", "", "The programming language that this project uses (python, java, javascript)")
	createProjectCmd.MarkFlagRequired("language")
	createProjectCmd.Flags().BoolVar(&createRepoFlag, "create-repo", false, "Create a github repository")
	createProjectCmd.Flags().StringVarP(&descriptionFlag, "description", "d", "", "A summary of what the project is about")

	return createProjectCmd
}