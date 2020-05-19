package cmd

import (
	"github.com/BooookStore/RedmineCLI/cmd/service"
	"github.com/BooookStore/RedmineCLI/cmd/writer"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show story and tasks",
	Run: func(cmd *cobra.Command, args []string) {
		client := service.NewClient("http://localhost:8080", "290046cc011a116826e9ce2c54705b58ba98aba1")
		broker := service.Broker{Client: client}
		issues, err := broker.GetIssues(args[0], args[1])
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
}
