package cmd

import (
	"github.com/BooookStore/RedmineCLI/cmd/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show story and tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client := service.NewClient(viper.GetString("redmine-url"), viper.GetString("redmine-api-key"))
		broker := &service.Broker{Client: client}
		writer := &service.Writer{Out: cmd.OutOrStdout()}

		inspectIssueId, err := cmd.Flags().GetInt("inspect")
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		if inspectIssueId != -1 {
			err = service.PrintIssue(broker, writer, viper.GetString("project"), viper.GetString("sprint"), inspectIssueId)
			if err != nil {
				cmd.PrintErr(err)
				return
			}
		} else {
			err = service.PrintIssues(broker, writer, viper.GetString("project"), viper.GetString("sprint"))
			if err != nil {
				cmd.PrintErr(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringP("project", "p", "", "project name")
	statusCmd.Flags().StringP("sprint", "s", "", "sprint name")
	statusCmd.Flags().IntP("inspect", "i", -1, "issue id")
	err := viper.BindPFlag("project", statusCmd.Flags().Lookup("project"))
	err = viper.BindPFlag("sprint", statusCmd.Flags().Lookup("sprint"))

	if err != nil {
		statusCmd.PrintErr(err)
		os.Exit(1)
	}
}
