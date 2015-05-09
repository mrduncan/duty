package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
)

// The main command runner which all commands are registered to.
var CmdRunner = NewRunner()

type Args struct {
	Command string
	Params  []string
}

func NewArgs(args []string) *Args {
	var command string
	var params []string

	if len(args) == 0 {
		params = []string{}
	} else {
		command = args[0]
		params = args[1:]
	}

	return &Args{Command: command, Params: params}
}

type Command struct {
	Name        string
	Description string
	Run         func(*Args) error
}

type Runner struct {
	commands map[string]*Command
}

func NewRunner() *Runner {
	return &Runner{commands: make(map[string]*Command)}
}

func (r *Runner) Register(cmd *Command) {
	r.commands[cmd.Name] = cmd
}

func (r *Runner) Run(name string, args *Args) error {
	if name == "" {
		r.commands["help"].Run(nil)
		return errors.New("")
	} else {
		cmd := r.commands[name]
		if cmd == nil {
			return fmt.Errorf("'%s' is not a duty command\nRun 'duty help' for usage.", name)
		}
		return cmd.Run(args)
	}
}

// All returns all registered commands sorted by name.
func (r *Runner) All() []*Command {
	all := make([]*Command, 0, len(r.commands))
	for _, cmd := range r.commands {
		all = append(all, cmd)
	}
	sort.Sort(byName(all))
	return all
}

func (r *Runner) Execute() int {
	args := NewArgs(os.Args[1:])
	err := r.Run(args.Command, args)
	if err != nil {
		message := err.Error()
		if message != "" {
			fmt.Println(message)
		}
		return 1
	}

	return 0
}

type byName []*Command

func (c byName) Len() int {
	return len(c)
}
func (c byName) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c byName) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}
