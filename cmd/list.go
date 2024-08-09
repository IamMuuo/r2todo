/*
Copyright © 2024 Erick Heart <hearteric57@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all currently saved todos",
	Long:  `Lists all tasks stored by the the program.`,
	Run: func(cmd *cobra.Command, args []string) {
		showComplete, showOverdue := false, false

		showAll, err := cmd.Flags().GetBool("show-all")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed parsing show all")
			os.Exit(2)

		}

		showComplete, err = cmd.Flags().GetBool("show-complete")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed parsing show-complete flag")
			os.Exit(2)

		}

		if showAll {
			showComplete = true
			showOverdue = true
		}

		todoController.ListTodos(showComplete, showOverdue)
	},
}

func init() {

	listCmd.Flags().BoolP("show-all", "e", false, "Specifies whether to show all todos")
	listCmd.Flags().BoolP("show-complete", "c", true, "Specifies whether to show all todos")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
