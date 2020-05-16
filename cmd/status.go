/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/BooookStore/RedmineCLI/cmd/service"
	"github.com/spf13/cobra"
	"log"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show story and tasks",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := service.NewClient("http://localhost:8080", "290046cc011a116826e9ce2c54705b58ba98aba1")
		if err != nil {
			log.Fatal(err)
		}
		broker := service.Broker{Client: client}
		issues, err := broker.GetIssues(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(issues)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
