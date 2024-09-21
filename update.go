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

import tea "github.com/charmbracelet/bubbletea"


// Note: This function is intended for internal use by the Bubble Tea framework
// and should not be called directly by the end user.
// Update is the internal update function required by the Bubble Tea framework.
// Update is the main update function for the UI, called on every input event.
// It processes the input based on the current question index and the user's
// keypresses. The method handles various key events like Enter, Arrow keys, and 
// special keys like Ctrl+C to quit the program.
//
// If all questions have been answered (i.e., u.index >= len(u.questions)), 
// it quits the interactive UI.
//
// Params:
//   msg tea.Msg: The message containing user input (typically key presses).
//
// Returns:
//   (tea.Model, tea.Cmd): The updated model (ui) and command to be executed.
//
// Behavior:
//   - tea.KeyCtrlC: Quits the interactive UI.
//   - tea.KeyEnter: Handles "Enter" key to process the current question.
//   - tea.KeyDown: Moves to the next option in the UI.
//   - tea.KeyUp: Moves to the previous option in the UI.
//   - tea.KeyRight: Handles "Right Arrow" key input.
//   - tea.KeyLeft: Handles "Left Arrow" key input.
//   - tea.KeyTab, tea.KeySpace, tea.KeyBackspace: Handles additional key
//     inputs for Tab, Space, and Backspace keys.
//   - default: Passes other keys to the handleDefaultKey function for further 
//     handling.

func (u *ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if u.index >= len(u.questions) {
		return u, tea.Quit
	}

	u.errorMsg = ""

	switch msgType := msg.(type) {

	case tea.KeyMsg:

		switch msgType.Type {

		case tea.KeyCtrlC:
			return u, tea.Quit
			
		case tea.KeyEnter:
			return u.handleEnter()
			
		case tea.KeyDown:
			return u.handleDown()
			
		case tea.KeyUp:
			return u.handleUp()
			
		case tea.KeyRight:
			return u.handleRight()
			
		case tea.KeyLeft:
			return u.handleLeft()
			
		case tea.KeyTab, tea.KeySpace, tea.KeyBackspace:
			return u.handleTabSpaceBackspace()
			
		default:
			return u.handleDefaultKey(msgType)
		}
	}

	return u, nil
}

func (u *ui) handleEnter() (tea.Model, tea.Cmd) {

	if u.index < len(u.questions) {

		if u.questions[u.index].FilePath && u.input != "" {
			return u.handleFilePathEnter()

		} else if len(u.questions[u.index].Options) > 0 {
			return u.handleOptionsEnter()

		} else {
			return u.handleInputEnter()
		}
	}

	return u, tea.Quit
}

func (u *ui) handleFilePathEnter() (tea.Model, tea.Cmd) {

	u.questionFilePath()

	if len(u.filesList) > 0 {

		selectedFile := u.filesList[u.cursor]

		err := u.questions[u.index].Handler(selectedFile)

		if err != nil {

			u.errorMsg = err.Error()
			u.input = ""
			u.cursor = 0

			return u, nil
		}

		u.userAnswers = append(u.userAnswers, u.input)
		u.input = ""
		u.index++
		u.cursor = 0

		if u.index >= len(u.questions) {
			return u, tea.Quit
		}

		return u, nil
	}

	u.errorMsg = "Directory does not exist or no files in directory"
	return u, nil
}

func (u *ui) handleOptionsEnter() (tea.Model, tea.Cmd) {

	err := u.questions[u.index].Handler(u.questions[u.index].Options[u.cursor])

	if err != nil {
		u.errorMsg = err.Error()
		u.input = ""
		u.cursor = 0

		return u, nil
	}

	u.userAnswers = append(u.userAnswers, u.input)
	u.input = ""
	u.index++
	u.cursor = 0

	if u.index >= len(u.questions) {
		return u, tea.Quit
	}

	return u, nil
}

func (u *ui) handleInputEnter() (tea.Model, tea.Cmd) {

	err := u.questions[u.index].Handler(u.input)

	if err != nil {
		u.errorMsg = err.Error()
		u.input = ""
		u.cursor = 0

		return u, nil
	}

	u.userAnswers = append(u.userAnswers, u.input)
	u.input = ""
	u.index++
	u.cursor = 0

	if u.index >= len(u.questions) {
		return u, tea.Quit
	}

	return u, nil
}

func (u *ui) handleDown() (tea.Model, tea.Cmd) {

	if u.questions[u.index].FilePath && len(u.filesList) > 0 {

		if u.cursor < len(u.filesList)-1 {
			u.cursor++
		}
		
	} else if len(u.questions[u.index].Options) > 0 {

		if u.cursor < len(u.questions[u.index].Options)-1 {
			u.cursor++
		}
	}

	return u, nil
}

func (u *ui) handleUp() (tea.Model, tea.Cmd) {

	if u.cursor > 0 {
		u.cursor--
	}

	return u, nil
}

func (u *ui) handleRight() (tea.Model, tea.Cmd) {

	if u.questions[u.index].FilePath {
		u.handleFilePathRight()
		
	} else if len(u.questions[u.index].Options) > 0 {
		u.handleOptionsRight()
	}

	return u, nil
}

func (u *ui) handleLeft() (tea.Model, tea.Cmd) {
	if u.cursor >= pageSize {
		u.cursor -= pageSize
	} else {
		u.cursor = 0
	}
	return u, nil
}

func (u *ui) handleTabSpaceBackspace() (tea.Model, tea.Cmd) {

	if len(u.input) > 0 {
		u.input = u.input[:len(u.input)-1]
	}

	return u, nil
}

func (u *ui) handleDefaultKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	u.input += msg.String()

	return u, nil
}

func (u *ui) handleFilePathRight() {

	nextPage := u.cursor + pageSize

	if nextPage >= len(u.filesList) {
		u.cursor = 0
	} else {
		u.cursor = nextPage
	}
}

func (u *ui) handleOptionsRight() {

	nextPage := (u.cursor/pageSize + 1) * pageSize

	if nextPage < len(u.questions[u.index].Options) {
		u.cursor = nextPage
	} else {
		u.cursor = len(u.questions[u.index].Options) - 1
	}
}
