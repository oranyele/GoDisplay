package display

import (
	"fmt"
	"io"
	"os"
	"time"
)

type StepStatus int

const (
	StepSuccess StepStatus = iota
	StepError
)

type TerminalDisplayService interface {
	Message(title string)
	StepMessage(msg string, status StepStatus)
	DisplayList(list []string)
	Finish()
}

type TerminalDisplay struct {
	timeStarted time.Time
	out         io.Writer
}

func NewTerminalDisplayService() TerminalDisplayService {
	return &TerminalDisplay{
		timeStarted: time.Now(),
		out:         os.Stdout,
	}
}

func (d *TerminalDisplay) Message(title string) {
	fmt.Fprintln(d.out, title)
}

func (d *TerminalDisplay) StepMessage(msg string, status StepStatus) {
	var symbol string
	var color string
	switch status {
	case StepSuccess:
		symbol = "[OK]"
		color = "\033[32m"
	case StepError:
		symbol = "[FAIL]"
		color = "\033[31m"
	}

	fmt.Fprintf(d.out, "\t%s %s%s\033[0m\n", msg, color, symbol)
}

func (d *TerminalDisplay) DisplayList(list []string) {
	for _, item := range list {
		fmt.Fprintf(d.out, "\t\t%s\n", item)
	}
}

func (d *TerminalDisplay) Finish() {
	fmt.Fprintf(d.out, "All steps completed - took %s\033[0m\n\033[32mDone!\033[0m\n", time.Since(d.timeStarted).Round(time.Second).String())
}
