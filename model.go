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
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	spfcbr "github.com/spf13/cobra"
)

const pageSize = 10

type ui struct {
	cobra       *spfcbr.Command
	questions   []Question
	index       int
	input       string
	cursor      int
	errorMsg    string
	userAnswers []string
	filesList   []string
}

// SetCobra links a Cobra command with the UI.
// Parameters:
// - cobra: The Cobra command instance to link with the UI.
func (u *ui) SetCobra(cobra *spfcbr.Command) {
	u.cobra = cobra
}

// Note: This function is intended for internal use by the Bubble Tea framework
// and should not be called directly by the end user.
// Init initializes the UI. 
// It returns nil as no specific initialization is needed.
func (u *ui) Init() tea.Cmd {
	return nil
}

// SetQuestions sets the list of questions to be presented during the UI interaction.
//
// Parameters:
// - questions: A slice of Question structs representing the questions to display.
func (u *ui) SetQuestions(questions []Question) {
	u.questions = questions
}

// RunInteractiveUI starts the interactive UI by running a Bubble Tea program.
// It initializes the index and cursor, then starts the UI loop.
func (u *ui) RunInteractiveUI() {
	u.index = 0
	u.cursor = 0

	p := tea.NewProgram(u)

	if _, err := p.Run(); err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}

// AfterPreRun sets the PreRun hook for Cobra to execute the interactive UI after the PreRun logic.
// If PreRun is already defined, it ensures the original logic is preserved.
func (u *ui) AfterPreRun() {

	if u.cobra == nil {
		log.Println("Cobra instance is not set")
	}

	existingPreRun := u.cobra.PreRun

	if existingPreRun != nil {
		u.cobra.PreRun = func(cmd *spfcbr.Command, args []string) {
		existingPreRun(cmd, args)
		u.RunInteractiveUI()
		}

	} else {
		u.cobra.PreRun = func(cmd *spfcbr.Command, args []string) {
		u.RunInteractiveUI()
		}
	}
}

// BeforePreRun sets the PreRun hook for Cobra to execute the interactive UI before the PreRun logic.
// If PreRun is already defined, it ensures the UI is executed first and the original logic is preserved.
func (u *ui) BeforePreRun() {

	if u.cobra == nil {
		log.Println("Cobra instance is not set")
	}

	existingPreRun := u.cobra.PreRun

	if existingPreRun != nil {
		u.cobra.PreRun = func(cmd *spfcbr.Command, args []string) {
		u.RunInteractiveUI()
		existingPreRun(cmd, args)
		}

	} else {
		u.cobra.PreRun = func(cmd *spfcbr.Command, args []string) {
		u.RunInteractiveUI()
		}
	}
}

// BeforeRun sets the Run hook for Cobra to execute the interactive UI before the command's main logic.
// If Run is already defined, it ensures the UI is executed first and preserves the original logic.
func (u *ui) BeforeRun() {

	if u.cobra == nil {
		log.Println("Cobra instance is not set")
	}

	existingRun := u.cobra.Run

	if existingRun != nil {
		u.cobra.Run = func(cmd *spfcbr.Command, args []string) {
		u.RunInteractiveUI()
		existingRun(cmd, args)
		}
		
	} else {
		u.cobra.Run = func(cmd *spfcbr.Command, args []string) {
		u.RunInteractiveUI()
		}
	}
}

// AfterRun sets the Run hook for Cobra to execute the interactive UI after the command's main logic.
// If Run is already defined, it ensures the original logic is preserved and the UI is run afterward.
func (u *ui) AfterRun() {

	if u.cobra == nil {
		log.Println("Cobra instance is not set")
	}

	existingRun := u.cobra.Run

	if existingRun != nil {
		u.cobra.Run = func(cmd *spfcbr.Command, args []string) {
		existingRun(cmd, args)
		u.RunInteractiveUI()
		}
		
	} else {
		u.cobra.Run = func(cmd *spfcbr.Command, args []string) {
		u.RunInteractiveUI()

		}
	}
}

func (u *ui) questionFilePath() {

	u.filesList = nil

	fullPath := ""

	if u.input == "." {
		fullPath, _ = filepath.Abs(u.input)

	} else {
		fullPath = u.input
	}

	if _, err := os.Stat(fullPath); err == nil {

		files, _ := filepath.Glob(filepath.Join(fullPath, "*"))

		for _, file := range files {

			fileInfo, _ := os.Stat(file)

			if fileInfo != nil && !fileInfo.IsDir() {
				u.filesList = append(u.filesList, file)
			}
		}

	} else {
		u.errorMsg = "Directory does not exist"
	}
}
