package executor

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"orion/internal/config"
	"orion/internal/plan"
)

type Logger struct {
	file *os.File
	out  io.Writer
}

func NewLogger() (*Logger, error) {
	logsDir := config.LogsDir()
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs dir: %w", err)
	}

	filename := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02_15-04-05"))
	path := filepath.Join(logsDir, filename)

	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %w", err)
	}

	// Tee output to both file and stdout
	multi := io.MultiWriter(os.Stdout, f)

	return &Logger{
		file: f,
		out:  multi,
	}, nil
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}

func (l *Logger) Write(p []byte) (n int, err error) {
	return l.out.Write(p)
}

func (l *Logger) Log(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("15:04:05")
	fmt.Fprintf(l.out, "[%s] %s\n", timestamp, msg)
}

func Execute(p *plan.CommandPlan, cfg config.Config) error {
	logger, err := NewLogger()
	if err != nil {
		fmt.Printf("⚠️  Failed to init logger: %v\n", err)
		// continue without file logging if it fails?
		// For now, let's treat it as fatal or fallback to stdout only
		logger = &Logger{out: os.Stdout}
	} else {
		defer logger.Close()
		fmt.Printf("📝 Logging to %s\n", logger.file.Name())
	}

	logger.Log("Executing plan: %s", p.Summary)
	if p.Cwd != "" {
		logger.Log("Working directory: %s", p.Cwd)
	}

	for i, cmd := range p.Commands {
		logger.Log("Step %d/%d: %s", i+1, len(p.Commands), cmd.Cmd)

		if err := runCommand(cmd.Cmd, p.Cwd, logger.out); err != nil {
			logger.Log("❌ Failed: %v", err)
			return fmt.Errorf("step %d failed: %w", i+1, err)
		}
	}

	logger.Log("✅ Plan execution completed successfully.")
	return nil
}

func runCommand(commandStr, cwd string, out io.Writer) error {
	// Simple shell execution for now
	cmd := exec.Command("sh", "-c", commandStr)
	if cwd != "" {
		cmd.Dir = cwd
	}

	cmd.Stdout = out
	cmd.Stderr = out // Merge stderr into stdout for the log

	return cmd.Run()
}
