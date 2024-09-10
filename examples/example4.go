package main

import (
	"fmt"

	cobra_ui "github.com/sabouaram/cobra-ui"
)

func main() {

	var (
		passwordEntered bool
		pwd             string
	)

	ui := cobra_ui.New()
	ui.SetQuestions([]cobra_ui.Question{
		{
			Text:         "Enter your password: ",
			PasswordType: true,
			Handler: func(password string) error {
				passwordEntered = true
				pwd = password
				return nil
			},
		},
	})
	ui.RunInteractiveUI()
	if passwordEntered {
		fmt.Printf("Password entered => %s\n", pwd)
	}

}
