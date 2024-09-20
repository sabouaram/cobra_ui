package main

import (
	"fmt"

	"github.com/fatih/color"
	cobra_ui "github.com/sabouaram/cobra_ui"
)

func main() {
	var selectedFile string
	ui := cobra_ui.New()
	ui.SetQuestions([]cobra_ui.Question{
		{
			Text:     "Select a file:",
			Color:    color.FgCyan,
			FilePath: true,
			Handler: func(filePath string) error {
				selectedFile = filePath
				return nil
			},
		},
	})
	ui.RunInteractiveUI()
	fmt.Printf("done")
}
