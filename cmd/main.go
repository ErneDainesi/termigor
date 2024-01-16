package main

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"github.com/ErneDainesi/termigor"
	"github.com/creack/pty"
)

func main() {
	os.Setenv("TERM", "dumb")
	c := exec.Command("/bin/bash")
	p, err := pty.Start(c) // start pseudoterminal

	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer c.Process.Kill()

	buffer := [][]rune{}
	reader := bufio.NewReader(p)

	// Goroutine that reads from pty
	go func() {
		line := []rune{}
		buffer = append(buffer, line)
		for {
			r, _, err := reader.ReadRune()

			if err != nil {
				if err == io.EOF {
					return
				}
				os.Exit(0)
			}

			line = append(line, r)
			buffer[len(buffer)-1] = line
			if r == '\n' {
				line = []rune{}
				buffer = append(buffer, line)
			}
		}
	}()

    ui := termigor.StartUI(p)

	// Goroutine that renders to UI
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			ui.SetText("")
			var lines string
			for _, line := range buffer {
				lines = lines + string(line)
			}
			ui.SetText(string(lines))
		}
	}()
}
