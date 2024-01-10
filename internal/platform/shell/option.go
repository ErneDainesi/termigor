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
	env   []string
}

func NewOptions(options ...Option) (Options, error) {
	opt := Options{
		shell: os.Getenv("SHELL"),
	}

	for _, op := range options {
		opt = op(opt)
	}

	if err := validateOpts(opt); err != nil {
		return Options{}, err
	}

	return opt, nil
}

func (opt Options) Shell() string {
	return opt.shell
}

func (opt Options) Env() []string {
	return opt.env
}

func WithShell(s string) Option {
	return func(options Options) Options {
		options.shell = s
		return options
	}
}

func WithEnv(envVar string) Option {
	return func(opts Options) Options {
		opts.env = append(opts.env, envVar)
		return opts
	}
}

func WithEnvVariables(evs []string) Option {
	return func(opts Options) Options {
		opts.env = evs
		return opts
	}
}

func validateOpts(options Options) error {
	if options.shell == "" {
		return ErrNoShell
	}
	return nil
}
