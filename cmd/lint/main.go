package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const defaultGoLintModule = "github.com/skulidropek/GoLint/cmd/go-lint@13befa895e15e52dbffec7bf827fa69bb5fcf364"

func main() {
	args := os.Args[1:]

	if bin, err := exec.LookPath("go-lint"); err == nil {
		exitIfError(runCommand(bin, args...))
		return
	}

	module := strings.TrimSpace(os.Getenv("DEVKIT_GO_LINT_MODULE"))
	if module == "" {
		module = defaultGoLintModule
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
	fmt.Fprintf(os.Stderr, "devkit lint wrapper failed: %v\n", err)
	os.Exit(1)
}
