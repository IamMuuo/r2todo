/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a todo item",
	Long: `The 'create' command creates a todo item. 

The id and date created are automatically created
`,
	Run: func(cmd *cobra.Command, args []string) {

		taskDescription, err := cmd.Flags().GetString("description")
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"Failed to read the --description flag which is needed while creating task\n",
			)
			os.Exit(1)
		}

		if taskDescription == "" {
			fmt.Fprintf(os.Stderr, "Please specify the todo's description\n")
			os.Exit(2)
		}

		if err := todoController.CreateTodo(taskDescription); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create todo with error\n\n%s", err.Error())
			os.Exit(2)
		}
	},
}

func init() {
	createCmd.Flags().StringP("description", "d", "", "A task's description")
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
