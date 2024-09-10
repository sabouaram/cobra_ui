package main

import (
	"errors"
	"fmt"
	"strconv"

	cobra_ui "github.com/sabouaram/cobra-ui"
)

func main() {

	var (
		age int
		err error
	)

	tui := cobra_ui.New()
	tui.SetQuestions([]cobra_ui.Question{
		{
			Text: "Enter your age: ",
			Handler: func(input string) error {
				age, err = strconv.Atoi(input)
				if err != nil {
					return errors.New("age must be an integer - retry")
				}
				return nil
			},
		},
	})
	tui.RunInteractiveUI()
	fmt.Printf("Your entered age is %d \n", age)
}
