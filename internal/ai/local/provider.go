package local

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Local struct {
	BinPath string
}

func (l Local) Complete(ctx context.Context, prompt string) (string, error) {
	if l.BinPath == "" {
		return "", fmt.Errorf("local AI binary path not configured")
	}

	// Safety: Command execution of the sidecar
	// We expect the sidecar to accept the prompt via STDIN and output JSON to STDOUT
	cmd := exec.CommandContext(ctx, l.BinPath)

	// Set up pipes
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = strings.NewReader(prompt)

	// Execute with timeout (implicit if ctx has deadline, but good to be explicit about expectation)
	// For local LLM invocation, we might need a generous timeout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("local AI execution failed: %v (stderr: %s)", err, stderr.String())
	}

	output := strings.TrimSpace(stdout.String())
	if output == "" {
		return "", fmt.Errorf("local AI returned empty response (stderr: %s)", stderr.String())
	}

	return output, nil
}

func (l Local) Health(ctx context.Context) error {
	if l.BinPath == "" {
		return fmt.Errorf("local AI binary path not configured")
	}

	info, err := os.Stat(l.BinPath)
	if err != nil {
		return fmt.Errorf("local AI binary not found at %s: %v", l.BinPath, err)
	}
	if info.IsDir() {
		return fmt.Errorf("local AI binary path is a directory: %s", l.BinPath)
	}
	if info.Mode()&0111 == 0 {
		return fmt.Errorf("local AI binary is not executable: %s", l.BinPath)
	}

	// deep health check: try running with --version or similar if supported
	// for now, existence is enough
	return nil
}
