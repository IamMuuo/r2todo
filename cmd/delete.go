/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a todo item from cache",
	Long: `Delete

The delete subcommand either deletes a specified todo
or deletes all if passed the --all hence restating
.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("task-id")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Please specify the todo ID to update completion status")
			os.Exit(2)
		}

		deleteAll, err := cmd.Flags().GetBool("delete-all")
		if id == 0 && deleteAll == false {
			fmt.Fprintf(os.Stderr, "Please specify the todo ID to delete or use the --delete-all flag\n")
			os.Exit(2)
		}

		if err = todoController.Delete(id, deleteAll); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update task completion status\nError: %s\n", err.Error())
			os.Exit(2)
		}
	},
}

func init() {

	deleteCmd.Flags().BoolP("delete-all", "a", false, "Specifies whether to delete all todos")
	deleteCmd.Flags().IntP("task-id", "i", 0, "A task's id in question to be manipulated")
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
