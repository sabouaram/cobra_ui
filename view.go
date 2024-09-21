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
	"fmt"
	"strings"

	"github.com/fatih/color"
)


// View renders the current state of the UI as a string.
// This method is part of the internal Bubble Tea framework lifecycle and
// is called automatically to display the UI. It should not be called
// directly by end users. 
func (u *ui) View() string {

	if u.index >= len(u.questions) {
		return ""
	}

	question := u.questions[u.index]

	view := question.Text

	if question.Color != 0 {
		colorFunc := color.New(question.Color).SprintFunc()
		view = colorFunc(view)
	}

	if question.FilePath && u.input != "" {

		view += u.input + "\n"

		if u.errorMsg != "" {
		view += "Error: " + u.errorMsg + "\n"
		}

		u.appendFilePathView(&view)

	} else if len(question.Options) > 0 {
		u.appendOptionsView(&view)

	} else {
		u.appendInputView(&view, question)
	}

	return view
}

func (u *ui) appendFilePathView(view *string) {

	u.questionFilePath()

	if len(u.filesList) > 0 {

		*view += "Files in folder:\n"

		totalPages := (len(u.filesList) + pageSize - 1) / pageSize

		currentPage := u.cursor/pageSize + 1

		start := (currentPage - 1) * pageSize

		end := start + pageSize

		if end > len(u.filesList) {
			end = len(u.filesList)
		}

		if start >= len(u.filesList) {
			u.cursor = 0
			currentPage = 1
			start = 0
			end = pageSize
		}

		for i := start; i < end; i++ {

			cursor := " "

			if i == u.cursor {

				if u.questions[u.index].CursorStr != "" {
					cursor = u.questions[u.index].CursorStr

				} else {
					cursor = "→"
				}
			}

			*view += fmt.Sprintf("%s %d. %s\n", cursor, i+1, u.filesList[i])
		}

		*view += fmt.Sprintf("\nPage %d/%d\n", currentPage, totalPages)

	} else {
		*view += "No files in folder\n"
		u.cursor = 0
	}
}

func (u *ui) appendOptionsView(view *string) {

	*view += "\n"

	totalOptions := len(u.questions[u.index].Options)

	totalPages := (totalOptions + pageSize - 1) / pageSize

	currentPage := u.cursor/pageSize + 1

	start := (currentPage - 1) * pageSize

	end := start + pageSize

	if end > totalOptions {
		end = totalOptions
	}

	if start >= totalOptions {
		u.cursor = 0
		currentPage = 1
		start = 0
		end = pageSize
	}

	for i := start; i < end; i++ {

		cursor := " "

		if i == u.cursor {

			if u.questions[u.index].CursorStr != "" {
				cursor = u.questions[u.index].CursorStr

			} else {
				cursor = "→"
			}
		}

		*view += fmt.Sprintf("%s %d. %s\n", cursor, i+1, u.questions[u.index].Options[i])
	}

	*view += fmt.Sprintf("\nPage %d/%d\n", currentPage, totalPages)
}

func (u *ui) appendInputView(view *string, question Question) {

	if len(u.input) > 0 {

		if question.PasswordType {
			*view += strings.Repeat("*", len(u.input)) + "\n"

		} else {
			*view += u.input + "\n"
		}

	} else {
		*view += "\n"
	}

	if u.errorMsg != "" {
		*view += "\nError: " + u.errorMsg + "\n"

	}
}
