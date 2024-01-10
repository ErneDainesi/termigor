package shell

import (
	"errors"
	"testing"
)

func TestOptions_NoShellError(t *testing.T) {
	opt := NewOptions(WithShell(""))
	_, err := opt.Shell()

	if !errors.Is(err, ErrNoShell) {
		t.Fatalf("expecting ErrNoShell, other error found: %s", err)
	}
}

func TestOptions_WithShell(t *testing.T) {
	anyShell := "/bin/zsh"
	opt := NewOptions(WithShell(anyShell))

	shell, err := opt.Shell()

	if err != nil {
		t.Fatalf("expected no error, found %s", err)
	}
	if shell != anyShell {
		t.Fatalf("expected shell %s, found %s", anyShell, shell)
	}
}
