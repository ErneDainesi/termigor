package shell

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

// Shell is the interface for command line processors.
type Shell interface {
	ExecCmd(cmd string) CommandOutput
	ExecCmdStream(cmd string) error
}

// maybe posix?
type LinuxShell struct {
	opts Options
}

func NewLinuxShell(opt Options) *LinuxShell {
	return &LinuxShell{opt}
}

// ExecCmd executes cmd as the shell command. Whole parsing is
// done by a shell binary, here we just invoke it.
// result should be like: bash -c 'cmd' (quoted version of cmd for safety).
// ExecCmd waits until the whole program ends, for a async/streaming version
// check ExecCmdStream.
func (ls *LinuxShell) ExecCmd(cmd string) (CommandOutput, error) {
	var (
		stdoutBs = make([]byte, 0, 1024) // minimum buffer size
		stderrBs = make([]byte, 0, 1024) // minimum buffer size

		stdoutBuff = bytes.NewBuffer(stdoutBs)
		stderrBuff = bytes.NewBuffer(stderrBs)

		exitCode int
	)

	shell := ls.opts.Shell()

	c := exec.Command(shell, "-c", cmd)

	c.Env = append(c.Env, ls.opts.Env()...)
	c.Stdout = stdoutBuff
	c.Stderr = stderrBuff

	if err := c.Run(); err != nil {
		var exitError *exec.ExitError
		if !errors.As(err, &exitError) {
			return CommandOutput{}, fmt.Errorf("error running command: %w", err)
		}
		exitCode = exitError.ExitCode()
	}
	if c.Err != nil {
		return CommandOutput{}, fmt.Errorf("error with process execution: %w", c.Err)
	}

	return CommandOutput{
		Stdout:   stdoutBuff,
		Stderr:   stderrBuff,
		ExitCode: exitCode,
	}, nil
}

// ExecCmdStream TODO: implement
func (ls *LinuxShell) ExecCmdStream(cmd string) error {
	panic("unimplemented")
}

type CommandOutput struct {
	Err      error
	Stdout   *bytes.Buffer
	Stderr   *bytes.Buffer
	ExitCode int
}
