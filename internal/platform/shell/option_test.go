package shell

import (
	"slices"
	"testing"
)

func TestOptions_WithShell(t *testing.T) {
	anyShell := "/bin/zsh"
	opt, err := NewOptions(WithShell(anyShell))
	if err != nil {
		t.Fatalf("no error expected, found: %s", err)
	}

	shell := opt.Shell()

	if shell != anyShell {
		t.Fatalf("expected shell %s, found %s", anyShell, shell)
	}
}

func TestOptions_WithEnv(t *testing.T) {
	opt, err := NewOptions(
		WithShell("shell"),
		WithEnv("FOO=BAR"))
	if err != nil {
		t.Fatalf("no error expected, found: %s", err)
	}
	if !slices.Contains(opt.env, "FOO=BAR") {
		t.Fatalf("env var FOO=BAR not found")
	}
}

func TestOptions_WithEnvVariables(t *testing.T) {
	envVars := []string{
		"FOO=BAR",
		"BAR=BAZ",
	}
	opt, err := NewOptions(
		WithShell("shell"),
		WithEnvVariables(envVars))
	if err != nil {
		t.Fatalf("no error expected, found: %s", err)
	}

	for _, ev := range envVars {
		if !slices.Contains(opt.env, ev) {
			t.Errorf("env variable %s not found", ev)
		}
	}
}

func TestOptions_ErrNoShell(t *testing.T) {
	_, err := NewOptions(WithShell(""))
	if err == nil {
		t.Fatalf("expected ErrNoShell, found nil")
	}
}
