// Creates a new text file with the current directory as its tag and copies stdin into the window. If past arguments
// executes the arguments as a command and uses that command's stdout instead.
package main

import (
	"9fans.net/go/acme"
	"fmt"
	"io"
	"os"
	"path"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	win, err := acme.New()
	if err != nil {
		return fmt.Errorf("new window: %w", err)
	}
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("lookup wd: %w", err)
	}
	err = win.Name(path.Clean(wd)+"/") // Trailing slash makes right click file names work
	if err != nil {
		return fmt.Errorf("set name: %w", err)
	}
	err = win.Ctl("nomenu")
	if err != nil {
		return fmt.Errorf("hide menu: %w", err)
	}

	dst := writer{
		win: win,
		f:   "data",
	}
	_, err = io.Copy(dst, os.Stdin)
	if err != nil {
		return fmt.Errorf("copy data: %w", err)
	}

	err = win.Ctl("clean")
	if err != nil {
		return fmt.Errorf("mark window clean: %w", err)
	}
	err = win.Addr("0,0")
	if err != nil {
		return fmt.Errorf("set addr: %w", err)
	}
	err = win.Ctl("dot=addr")
	if err != nil {
		return fmt.Errorf("scroll to top: %w", err)
	}
	err = win.Ctl("show")
	if err != nil {
		return fmt.Errorf("show: %w", err)
	}
	return nil
}

// Writes to the given file. Implements io.Writer
type writer struct {
	win *acme.Win
	f string
}

func (w writer) Write(p []byte) (n int, err error) {
	return w.win.Write(w.f, p)
}

