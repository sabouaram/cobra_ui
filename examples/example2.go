package main

import (
	"fmt"

	"github.com/fatih/color"
	cobra_ui "github.com/sabouaram/cobra-ui"
)

func main() {
	var choice string
	ui := cobra_ui.New()
	ui.SetQuestions([]cobra_ui.Question{
		{
			CursorStr: "==>",
			Color:     color.FgCyan,
			Text:      "What is your preferred programming language?",
			Options:   []string{"Go", "Python", "JavaScript", "Java"},
			Handler: func(input string) error {
				choice = input
				return nil
			},
		},
	})
	ui.RunInteractiveUI()
	fmt.Printf("Selected choice: %s\n", choice)
}
