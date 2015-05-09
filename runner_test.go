package main

import (
	"errors"
	"github.com/bmizerany/assert"
	"testing"
)

func TestNewArgs(t *testing.T) {
	args := NewArgs([]string{})
	assert.Equal(t, "", args.Command)
	assert.Equal(t, 0, len(args.Params))

	args = NewArgs([]string{"command"})
	assert.Equal(t, "command", args.Command)
	assert.Equal(t, 0, len(args.Params))

	args = NewArgs([]string{"command", "param"})
	assert.Equal(t, "command", args.Command)
	assert.Equal(t, 1, len(args.Params))
}

func TestRunnerRun(t *testing.T) {
	runner := NewRunner()

	assert.NotEqual(t, nil, runner.Run("unknown", NewArgs([]string{})))

	runError := errors.New("known error")
	knownCmd := &Command{
		Name: "known",
		Run: func(args *Args) error {
			return runError
		},
	}
	runner.Register(knownCmd)
	assert.Equal(t, runError, runner.Run("known", NewArgs([]string{})))
}

func TestRunnerRunBlank(t *testing.T) {
	runner := NewRunner()

	var helpRun = false
	helpCmd := &Command{
		Name: "help",
		Run: func(args *Args) error {
			helpRun = true
			return nil
		},
	}
	runner.Register(helpCmd)
	runner.Run("", NewArgs([]string{}))
	assert.T(t, helpRun)
}

func TestRunnerAll(t *testing.T) {
	runner := NewRunner()
	assert.Equal(t, 0, len(runner.All()))

	// All returns an array sorted by name
	cmd1 := &Command{Name: "1"}
	cmd2 := &Command{Name: "2"}
	cmd3 := &Command{Name: "3"}
	runner.Register(cmd2)
	runner.Register(cmd1)
	runner.Register(cmd3)
	expected := []*Command{cmd1, cmd2, cmd3}
	assert.Equal(t, expected, runner.All())
}
