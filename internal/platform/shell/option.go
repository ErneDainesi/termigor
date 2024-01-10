package shell

import (
	"errors"
	"os"
)

var (
	ErrNoShell = errors.New("error: no shell found")
)

type Option func(options Options) Options

type Options struct {
	shell string
}

func NewOptions(options ...Option) Options {
	opt := Options{
		shell: os.Getenv("SHELL"),
	}

	for _, op := range options {
		opt = op(opt)
	}

	return opt
}

func (opt Options) Shell() (string, error) {
	if opt.shell == "" {
		return "", ErrNoShell
	}
	return opt.shell, nil
}

func WithShell(s string) Option {
	return func(options Options) Options {
		options.shell = s
		return options
	}
}
