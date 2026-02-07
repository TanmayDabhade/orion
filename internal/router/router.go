package router

import (
	"fmt"
	"net/url"
	"strings"

	"orion/internal/apps"
	"orion/internal/shortcuts"
	"orion/models"
)

type Result struct {
	Intent         models.Intent
	FallbackSearch bool
}

func Route(input string, shortcutMap map[string]string) (Result, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return Result{}, fmt.Errorf("input required")
	}

	if cmd, ok := shortcuts.Resolve(shortcutMap, trimmed); ok {
		return Result{
			Intent: models.Intent{
				Action: models.ActionRunShell,
				Args:   map[string]string{"command": cmd},
			},
		}, nil
	}

	lower := strings.ToLower(trimmed)
	if strings.HasPrefix(lower, "open ") {
		app := strings.TrimSpace(trimmed[5:])
		if app != "" {
			return Result{
				Intent: models.Intent{
					Action: models.ActionOpenApp,
					Args:   map[string]string{"app": app},
				},
			}, nil
		}
	}

	// Dynamic App Detection
	// If the input doesn't look like a command or URL, check if it's an app
	if path, ok := apps.Find(trimmed); ok {
		return Result{
			Intent: models.Intent{
				Action: models.ActionOpenApp,
				Args:   map[string]string{"app": path}, // Pass full path
			},
		}, nil
	}

	// Also check "app name" (e.g. "Google Chrome")
	if path, ok := apps.Find(strings.ReplaceAll(trimmed, " ", "")); ok {
		return Result{
			Intent: models.Intent{
				Action: models.ActionOpenApp,
				Args:   map[string]string{"app": path},
			},
		}, nil
	}

	if strings.HasPrefix(lower, "search ") {
		q := strings.TrimSpace(trimmed[7:])
		if q != "" {
			return Result{
				Intent: models.Intent{
					Action: models.ActionSearch,
					Args:   map[string]string{"query": q},
				},
			}, nil
		}
	}

	if isURL(trimmed) {
		return Result{
			Intent: models.Intent{
				Action: models.ActionOpenURL,
				Args:   map[string]string{"url": trimmed},
			},
		}, nil
	}

	if IsDomain(trimmed) {
		return Result{
			Intent: models.Intent{
				Action: models.ActionOpenURL,
				Args:   map[string]string{"url": "https://" + trimmed},
			},
		}, nil
	}

	return Result{
		Intent: models.Intent{
			Action: models.ActionSearch,
			Args:   map[string]string{"query": trimmed},
		},
		FallbackSearch: true,
	}, nil
}

func isURL(input string) bool {
	parsed, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func IsDomain(input string) bool {
	value := strings.TrimSpace(input)
	if value == "" || strings.ContainsAny(value, " /") {
		return false
	}
	if strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") {
		return false
	}
	parts := strings.Split(value, ".")
	if len(parts) < 2 {
		return false
	}
	allowed := map[string]bool{"com": true, "org": true, "edu": true}
	tld := strings.ToLower(parts[len(parts)-1])
	return allowed[tld]
}
