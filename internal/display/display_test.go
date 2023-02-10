package display

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTerminalProgressDisplay_NotNil(t *testing.T) {
	assert.NotNil(t, NewTerminalDisplayService())
}

func TestTerminalProgressDisplay_Message(t *testing.T) {
	buf := &bytes.Buffer{}

	display := &TerminalDisplay{
		out: buf,
	}

	mockMsg := "Message mocked"
	display.Message(mockMsg)
	expected := fmt.Sprintf("%s\n", mockMsg)
	assert.Equal(t, expected, buf.String())
}

func TestTerminalProgressDisplay_Update(t *testing.T) {
	buf := &bytes.Buffer{}
	display := &TerminalDisplay{out: buf}

	display.StepMessage("Step 1", StepSuccess)
	expected := "\tStep 1 \033[32m[OK]\033[0m\n"
	assert.Equal(t, expected, buf.String())
	buf.Reset()

	display.StepMessage("Step 2", StepError)
	expected = "\tStep 2 \033[31m[FAIL]\033[0m\n"
	assert.Equal(t, expected, buf.String())
	buf.Reset()
}

func TestTerminalProgressDisplay_DisplayList(t *testing.T) {
	buf := &bytes.Buffer{}

	display := &TerminalDisplay{
		out: buf,
	}

	mockList := []string{"item 1", "item 2", "item 3"}

	display.DisplayList(mockList)
	expected := "\t\titem 1\n\t\titem 2\n\t\titem 3\n"
	assert.Equal(t, expected, buf.String())
}

func TestTerminalProgressDisplay_Finish(t *testing.T) {
	buf := &bytes.Buffer{}
	mockTime := time.Now()
	display := &TerminalDisplay{
		out:         buf,
		timeStarted: mockTime,
	}

	time.Sleep(1000)
	display.Finish()
	expected := fmt.Sprintf("All steps completed - took %s\033[0m\n\033[32mDone!\033[0m\n", time.Since(mockTime).Round(time.Second).String())
	assert.Equal(t, expected, buf.String())
}
