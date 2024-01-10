package shell

import (
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

func TestOptions_ErrNoShell(t *testing.T) {
	_, err := NewOptions(WithShell(""))
	if err == nil {
		t.Fatalf("expected ErrNoShell, found nil")
	}
}
