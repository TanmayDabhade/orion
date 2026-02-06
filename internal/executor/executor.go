package executor

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"orion/internal/config"
	"orion/models"
)

func Execute(intent models.Intent, cfg config.Config) error {
	switch intent.Action {
	case models.ActionOpenURL:
		return runOpen(intent.Args["url"])
	case models.ActionOpenApp:
		app := intent.Args["app"]
		if app == "" {
			return fmt.Errorf("app name required")
		}
		return runCommand("open", "-a", app)
	case models.ActionSearch:
		query := intent.Args["query"]
		if query == "" {
			return fmt.Errorf("search query required")
		}
		engine := cfg.SearchEngine
		if engine == "" {
			engine = config.Default().SearchEngine
		}
		url := fmt.Sprintf(engine, url.QueryEscape(query))
		return runOpen(url)
	case models.ActionRunShell:
		cmd := intent.Args["command"]
		if cmd == "" {
			return fmt.Errorf("command required")
		}
		return runShell(cmd)
	default:
		return fmt.Errorf("unsupported action: %s", intent.Action)
	}
}

func runOpen(target string) error {
	if target == "" {
		return fmt.Errorf("url required")
	}
	return runCommand("open", target)
}

func runShell(command string) error {
	// Enhancement: If command ends in .app, assume it's a macOS app and use open -a
	if strings.HasSuffix(command, ".app") {
		return runCommand("open", "-a", command)
	}
	return runCommand("sh", "-c", command)
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
