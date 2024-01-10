package shell

import (
	"fmt"
	"os/exec"
)

// Shell is the interface for command line processors.
// It runs a command string from cmd an returns the lines
// read from stdout and stderr as "program output"
type Shell interface {
	ExecCmd(cmd string) CommandOutput
}

type LinuxShell struct {
	opts Options
}

func (ls *LinuxShell) ExecCmd(cmd string) (CommandOutput, error) {
	shell := ls.opts.Shell()
	c := exec.Command(shell, "-c")
	c.Env = append(c.Env, ls.opts.Env()...)

	if err := c.Run(); err != nil {
		return CommandOutput{}, fmt.Errorf("error running command: %w", err)
	}

	//TODO: add c.Stdout and c.Stderr as custom io.Writer to collect their output
	return CommandOutput{}, nil
}

type CommandOutput struct {
	err        error
	lines      []string
	statusCode uint
}

func (co CommandOutput) Err() error {
	return co.err
}

func (co CommandOutput) Lines() []string {
	return co.lines
}

func (co CommandOutput) StatusCode() uint {
	return co.statusCode
}
