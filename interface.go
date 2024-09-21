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

type Question struct {
	Text         string
	Options      []string
	FilePath     bool
	PasswordType bool
	Handler      func(string) error
	Color        color.Attribute
	CursorStr    string
}

type UI interface {
	SetQuestions(questions []Question)
	RunInteractiveUI()
	SetCobra(cobra *cobra.Command)
	AfterPreRun()
	BeforePreRun()
	AfterRun()
	BeforeRun()
}

func New() UI {
	return &ui{}
}
