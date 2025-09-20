package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const defaultGoTraceModule = "github.com/skulidropek/gotrace/cmd/gotrace-instrument@afda3736f26d21bb0ad41a341e96bee5990f7822"

func main() {
	args := os.Args[1:]

	if bin, err := exec.LookPath("gotrace-instrument"); err == nil {
		exitIfError(runCommand(bin, args...))
		return
	}

	module := strings.TrimSpace(os.Getenv("DEVKIT_GOTRACE_MODULE"))
	if module == "" {
		module = defaultGoTraceModule
	}

	cmdArgs := append([]string{"run", module}, args...)
	exitIfError(runCommand("go", cmdArgs...))
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	return cmd.Run()
}

func exitIfError(err error) {
	if err == nil {
		return
	}
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		os.Exit(exitErr.ExitCode())
	}
	fmt.Fprintf(os.Stderr, "devkit gotrace wrapper failed: %v\n", err)
	os.Exit(1)
}
