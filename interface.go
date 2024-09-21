/*
 *  MIT License
 *
 *  Copyright (c) 2024 Salim BOU ARAM
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 *
 */

package cobra_ui

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Question represents an interactive question that can be asked in the UI.
// It supports various input types like text, file selection, password entry, 
// and single-choice options.
//
// Fields:
// - Text: The question to be displayed to the user.
// - Options: A list of available options for single-choice questions. If empty, 
//   the user is expected to provide a text input.
// - FilePath: If true, the question prompts the user to select a file from a directory 
//   (with built-in pagination support).
// - PasswordType: If true, the input will be hidden for security, typically used 
//   for password entries.
// - Handler: A callback function that handles the user's input. It receives the 
//   input string and returns an error if the input is invalid or if there's an issue
//   processing the input.
// - Color: Defines the text color of the question. Uses `color.Attribute` from the 
//   `github.com/fatih/color` package.
// - CursorStr: The string used as a cursor or pointer in the UI when navigating 
//   through options. If not specified, the default cursor (`->`) is used.
type Question struct {
	Text         string
	Options      []string
	FilePath     bool
	PasswordType bool
	Handler      func(string) error
	Color        color.Attribute
	CursorStr    string
}

// UI defines the methods required to handle an interactive user interface (UI) 
// for terminal applications. It provides functionalities for setting up questions, 
// running the UI, and integrating with Cobra CLI commands.
//
// Methods:
// - SetQuestions: Allows setting a list of questions to be presented to the user. 
//   Each question can prompt the user for different types of inputs, such as file selection, 
//   text input, or multiple-choice options.
// - RunInteractiveUI: Starts the interactive UI session, displaying the configured 
//   questions to the user and collecting their responses.
// - SetCobra: Integrates the UI with a Cobra CLI command. This method links the 
//   interactive UI with a Cobra command, enabling the UI to run within the Cobra CLI's 
//   lifecycle.
// - AfterPreRun: This method is called after Cobra's PreRun phase, allowing for any 
//   additional setup or initialization that should occur after the PreRun function of a 
//   Cobra command.
// - BeforePreRun: This method is called before Cobra's PreRun phase, allowing the UI 
//   to perform setup tasks or make adjustments before the Cobra command's PreRun is executed.
// - AfterRun: Called after Cobra's Run phase, providing an opportunity for any post-execution 
//   steps or cleanup.
// - BeforeRun: Called before Cobra's Run phase, enabling setup or preparation tasks before 
//   the Cobra command is executed.
type UI interface {

	// SetQuestions sets the list of questions to be presented during the UI interaction.
	// Parameters:
	// - questions: A slice of Question structs representing the questions to display.
	SetQuestions(questions []Question)
	
	// RunInteractiveUI starts the interactive UI by running a Bubble Tea program.
	// It initializes the index and cursor, then starts the UI loop.
	RunInteractiveUI()
	
	// SetCobra links a Cobra command with the UI.
	// Parameters:
	// - cobra: The Cobra command instance to link with the UI.
	SetCobra(cobra *cobra.Command)

	// AfterPreRun sets the PreRun hook for Cobra to execute the interactive UI after the PreRun logic.
	// If PreRun is already defined, it ensures the original logic is preserved.
	AfterPreRun()

	// BeforePreRun sets the PreRun hook for Cobra to execute the interactive UI before the PreRun logic.
	// If PreRun is already defined, it ensures the UI is executed first and the original logic is preserved.
	BeforePreRun()

	// AfterRun sets the Run hook for Cobra to execute the interactive UI after the command's main logic.
	// If Run is already defined, it ensures the original logic is preserved and the UI is run afterward.
	AfterRun()

	// BeforeRun sets the Run hook for Cobra to execute the interactive UI before the command's main logic.
	// If Run is already defined, it ensures the UI is executed first and preserves the original logic.
	BeforeRun()
}
	
// New creates a new instance of the UI interface. It returns a concrete 
// implementation of the UI, which can be used to set up and run an interactive 
// user interface in terminal applications.
//
// The returned UI allows developers to configure questions, run the UI, and 
// integrate it with Cobra commands.
//
// Example usage:
//
//  ui := cobra_ui.New()
//  ui.SetQuestions([]cobra_ui.Question{
//      {
//          Text: "What is your preferred programming language?",
//          Options: []string{"Go", "Python", "JavaScript", "Java"},
//          Handler: func(input string) error {
//              fmt.Printf("Selected language: %s\n", input)
//              return nil
//          },
//      },
//  })
//  ui.RunInteractiveUI()
//
// Returns:
// - UI: A new instance of the UI interface, which can be configured and run.
func New() UI {
	return &ui{}
}
