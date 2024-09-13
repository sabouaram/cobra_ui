package main

import (
	"fmt"

	"github.com/sabouaram/cobra_ui"
	"github.com/spf13/cobra"
)

var choice string
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A sample cobra cli app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my cobra cli app!")
		fmt.Println("Thank you for choosing ", choice)
	},
}

func main() {

	ui := cobra_ui.New()
	ui.SetQuestions([]cobra_ui.Question{
		{
			Text:    "What is your preferred programming language?",
			Options: []string{"Go", "Python", "JavaScript"},
			Handler: func(input string) error {
				choice = input
				return nil
			},
		},
	})

	ui.SetCobra(rootCmd)
	ui.BeforeRun()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
