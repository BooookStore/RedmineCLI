package cmd

import (
	"github.com/BooookStore/RedmineCLI/cmd/service"
	"github.com/BooookStore/RedmineCLI/cmd/writer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show story and tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := viper.GetString("project")
		storyName := viper.GetString("story")

		client := service.NewClient("http://localhost:8080", "290046cc011a116826e9ce2c54705b58ba98aba1")
		broker := service.Broker{Client: client}
		issues, err := broker.GetIssues(projectName, storyName)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		w := &writer.Writer{Out: cmd.OutOrStdout()}
		err = w.PrintStories(issues.Issues...)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringP("project", "p", "", "project name")
	statusCmd.Flags().StringP("story", "s", "", "story name")
	err := viper.BindPFlag("story", statusCmd.Flags().Lookup("story"))
	err = viper.BindPFlag("project", statusCmd.Flags().Lookup("project"))

	if err != nil {
		statusCmd.PrintErr(err)
		os.Exit(1)
	}
}
