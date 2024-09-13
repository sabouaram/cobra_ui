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
	LastMessage string
	filesList   []string
}

func (u *ui) SetCobra(cobra *spfcbr.Command) {

	u.cobra = cobra
}

func (u *ui) Init() tea.Cmd {

	return nil
}

func (u *ui) SetQuestions(questions []Question) {

	u.questions = questions
}

func (u *ui) RunInteractiveUI() {

	u.index = 0
	u.cursor = 0

	p := tea.NewProgram(u)

	if _, err := p.Run(); err != nil {

		log.Println("Error:", err)

		os.Exit(1)
	}
}

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
