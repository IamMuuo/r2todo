/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Toggles the completion status of a todo to either true or false",
	Long: `Complete

The complete command toggles the completion status of a todo to either 
true or false depending on the current status of the todo. i.e 

If the task was set to true its set to false and vice versa
.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("task-id")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Please specify the todo ID to update completion status")
			os.Exit(2)
		}
		if id == 0 {
			fmt.Fprintf(os.Stderr, "Please specify the todo ID to update completion status")
			os.Exit(2)
		}

		if err = todoController.ToggleTodoStatus(id); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update task completion status\nError: %s\n", err.Error())
			os.Exit(2)
		}
	},
}

func init() {
	completeCmd.Flags().IntP("task-id", "i", 0, "A task's id in question to be manipulated")
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
