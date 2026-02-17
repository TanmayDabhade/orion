package plan

import (
	"fmt"
	"net/url"
	"orion/internal/config"
	"orion/models"
)

// FromIntent converts a legacy Intent models to a CommandPlan.
func FromIntent(intent models.Intent, cfg config.Config) *CommandPlan {
	var cmd string
	switch intent.Action {
	case models.ActionOpenURL:
		cmd = fmt.Sprintf("open '%s'", intent.Args["url"])
	case models.ActionOpenApp:
		cmd = fmt.Sprintf("open -a '%s'", intent.Args["app"])
	case models.ActionSearch:
		engine := cfg.SearchEngine
		if engine == "" {
			engine = "https://google.com/search?q=%s"
		}
		query := url.QueryEscape(intent.Args["query"])
		target := fmt.Sprintf(engine, query)
		cmd = fmt.Sprintf("open '%s'", target)
	case models.ActionRunShell:
		cmd = intent.Args["command"]
	default:
		cmd = fmt.Sprintf("echo 'Unknown action: %s'", intent.Action)
	}

	risk := RiskLow
	switch intent.Risk {
	case models.RiskHigh:
		risk = RiskHigh
	case models.RiskMedium:
		risk = RiskMedium
	}

	return &CommandPlan{
		Intent:  string(intent.Action),
		Summary: fmt.Sprintf("Execute %s", intent.Action),
		Cwd:     "",
		Commands: []Command{
			{Cmd: cmd, Risk: risk},
		},
	}
}
