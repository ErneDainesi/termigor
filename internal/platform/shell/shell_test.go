package shell

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestLinuxShell_ExecCmdShell(t *testing.T) {
	opts, err := NewOptions(WithShell(os.Getenv("SHELL")), WithEnv("FOO=BAR"))
	if err != nil {
		t.Fatalf("unexpected error creating options: %s", err)
	}
	ls := &LinuxShell{opts: opts}

	cd, err := ls.ExecCmd("echo $FOO")
	if err != nil {
		t.Fatal(err)
	}
	bs, err := io.ReadAll(cd.Stdout)
	if err != nil {
		t.Fatal(err)
	}
	var stb strings.Builder
	stb.Write(bs)

	if strings.Trim(stb.String(), "\n ") != "BAR" {
		t.Fatalf("unexpected output: %s", stb.String())
	}

	if cd.ExitCode != 0 {
		t.Fatalf("expected status code 0, found %d", cd.ExitCode)
	}

	if cd.Stderr.Len() != 0 {
		t.Fatalf("expected empty buffer, found buffer of length: %d", cd.Stderr.Len())
	}
}
