package gitconfd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"
)

// Execute looks into the provided directory and executes all scripts lying in
// there
func Execute(baseDirectory, hookType string) bool {
	confDirectory := path.Join(baseDirectory, fmt.Sprintf(".%s.d", hookType))
	if s, err := os.Stat(confDirectory); os.IsNotExist(err) || !s.IsDir() {
		return true
	} else {
		err := filepath.Walk(confDirectory, visit)
		if err != nil {
			return false
		}
	}
	return true
}

func visit(path string, f os.FileInfo, err error) error {
	var e error

	if !f.IsDir() {
		cmd := exec.Command(path)

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if e = cmd.Start(); e != nil {
			return fmt.Errorf("error: unable to start [%s]: %w", path, err)
		}
		if e = cmd.Wait(); e != nil {
			handleExecuteError(stdout, stderr, e, path)
		}
	}
	return e
}

// handleExecuteError does the lifting for an execute error
func handleExecuteError(stdout, stderr bytes.Buffer, err error, path string) bool {
	var (
		exitError *exec.ExitError
		ok        bool
	)

	if exitError, ok = err.(*exec.ExitError); !ok {
		fmt.Println(fmt.Sprintf("error: unable get exit code [%s]: %s", path, err.Error()))
	}
	printOutput(stdout, stderr)

	if _, ok := exitError.Sys().(syscall.WaitStatus); ok {
		fmt.Println(fmt.Sprintf("execution of [%s] failed", path))
	}
	return false
}

// printOutput dumps command output
func printOutput(stdout, stderr bytes.Buffer) {
	fmt.Println("output:")
	fmt.Println()
	fmt.Println(stdout.String())

	fmt.Println("error:")
	fmt.Println()
	fmt.Fprintln(os.Stderr, stderr.String())
}
