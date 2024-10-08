/*
Copyright © 2024 Erick Muuo hearteric57@gmail.com

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
	"os"

	"github.com/iammuuo/r2todo/configs"
	"github.com/iammuuo/r2todo/internal/controllers"
	"github.com/spf13/cobra"
)

var todoController controllers.TodoController

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "r2todo",
	Short: "A command-line task manager that does its task and does it well",
	Long: `r2todo is a command line todo application.

Its simple and quick to operate.. free from all internet
based distructions.

Your todo applications are also stored in a csv file.
You own your data
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.r2todo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("description", "d", "", "A task's description")
	rootCmd.Flags().IntP("task-id", "i", 0, "A task's description")
	rootCmd.Flags().BoolP("delete-all", "a", false, "Specifies whether to delete all todos")

	// Initialize the app's configuration
	cfg := configs.Config{}
	configs.LoadConfig(&cfg)

	todoController = controllers.TodoController{Cfg: &cfg}
}
