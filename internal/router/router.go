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

	// Check for "[shortcut] [app]" pattern (e.g. "gh chrome")
	parts := strings.Fields(trimmed)
	if len(parts) >= 2 {
		shortcutCmd, isShortcut := shortcuts.Resolve(shortcutMap, parts[0])
		if isShortcut && strings.HasPrefix(shortcutCmd, "open http") {
			// Check if the rest of the input identifies an app
			appName := strings.Join(parts[1:], " ")
			appPath := ""
			foundApp := false

			// 1. Check direct app detection
			if path, ok := apps.Find(appName); ok {
				appPath = path
				foundApp = true
			} else if path, ok := apps.Find(strings.ReplaceAll(appName, " ", "")); ok {
				appPath = path
				foundApp = true
			}

			// 2. Check setup aliases (shortcuts map)
			if !foundApp {
				if cmd, ok := shortcuts.Resolve(shortcutMap, appName); ok {
					// Expecting cmd like: open -a '/Applications/Google Chrome.app'
					if strings.HasPrefix(cmd, "open -a '") {
						// Extract path
						start := 9
						end := strings.LastIndex(cmd, "'")
						if end > start {
							appPath = cmd[start:end]
							foundApp = true
						}
					}
				}
			}

			if foundApp {
				// Construct new command: open -a "App Path" URL
				url := strings.TrimPrefix(shortcutCmd, "open ")
				newCmd := fmt.Sprintf("open -a '%s' %s", appPath, url)

				return Result{
					Intent: models.Intent{
						Action: models.ActionRunShell,
						Args:   map[string]string{"command": newCmd},
					},
				}, nil
			}
		}
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

	// Handle explicit "search" command
	if strings.HasPrefix(lower, "search ") {
		q := strings.TrimSpace(trimmed[7:])
		
		// Remove common prepositions for cleaner queries
		// "search for X" -> "X"
		// "search about X" -> "X"
		lowerQ := strings.ToLower(q)
		if strings.HasPrefix(lowerQ, "for ") {
			q = strings.TrimSpace(q[4:])
		} else if strings.HasPrefix(lowerQ, "about ") {
			q = strings.TrimSpace(q[6:])
		}
		
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
