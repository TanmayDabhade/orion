package cmd

import (
	"fmt"
	"os"

	"orion/internal/apps"
	"orion/internal/config"
	"orion/internal/shortcuts"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Index installed applications and create shortcuts",
	Long: `Scans your system for installed applications (.app bundles) and
creates shortcuts for them. This allows Orion to quickly open any
application by name.

Run this command after installation or whenever you install new apps.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔍 Scanning for installed applications...")

		// Get all applications from the system
		appMap := apps.List()

		if len(appMap) == 0 {
			fmt.Println("⚠️  No applications found in standard directories.")
			return nil
		}

		// Load existing shortcuts to merge with
		existingShortcuts, err := shortcuts.Load(config.ShortcutsPath())
		if err != nil {
			return fmt.Errorf("failed to load existing shortcuts: %w", err)
		}

		// Default web shortcuts
		defaults := map[string]string{
			"gh":       "open https://github.com",
			"so":       "open https://stackoverflow.com",
			"chat":     "open https://chatgpt.com",
			"claude":   "open https://claude.ai",
			"gemini":   "open https://aistudio.google.com",
			"gmail":    "open https://mail.google.com",
			"gcal":     "open https://calendar.google.com",
			"gmaps":    "open https://maps.google.com",
			"gdrive":   "open https://drive.google.com",
			"youtube":  "open https://youtube.com",
			"meet":     "open https://meet.google.com",
			"news":     "open https://news.ycombinator.com",
			"reddit":   "open https://reddit.com",
			"twitter":  "open https://twitter.com",
			"linkedin": "open https://linkedin.com",
			"local":    "open http://localhost:3000",
			"console":  "open https://console.cloud.google.com",
			"aws":      "open https://console.aws.amazon.com",
		}

		// Merge defaults: only add if key doesn't exist
		for key, cmd := range defaults {
			if _, exists := existingShortcuts[key]; !exists {
				existingShortcuts[key] = cmd
			}
		}

		// Merge: app shortcuts take lower priority than user-defined ones and defaults
		// Also add smart aliases for common apps
		aliases := map[string][]string{
			"google chrome":        {"chrome"},
			"google chrome canary": {"canary"},
			"firefox":              {"ff"},
			"visual studio code":   {"code", "vscode"},
			"sublime text":         {"subl"},
			"adobe photoshop 2024": {"photoshop", "ps"},
			"iterm":                {"iterm"},
			"warp":                 {"warp"},
			"brave browser":        {"brave"},
			"microsoft edge":       {"edge"},
			"arc":                  {"arc"},
		}

		for key, path := range appMap {
			if _, exists := existingShortcuts[key]; !exists {
				existingShortcuts[key] = fmt.Sprintf("open -a '%s'", path)
			}
			// Check for aliases
			if shortNames, ok := aliases[key]; ok {
				for _, short := range shortNames {
					if _, exists := existingShortcuts[short]; !exists {
						fmt.Printf("   ➕ Added alias: %s -> %s\n", short, key)
						existingShortcuts[short] = fmt.Sprintf("open -a '%s'", path)
					}
				}
			}
		}

		// Ensure config directory exists
		if err := os.MkdirAll(config.Dir(), 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}

		// Save the merged shortcuts
		if err := shortcuts.Save(config.ShortcutsPath(), existingShortcuts); err != nil {
			return fmt.Errorf("failed to save shortcuts: %w", err)
		}

		fmt.Printf("✅ Indexed %d applications to %s\n", len(appMap), config.ShortcutsPath())
		fmt.Println("💡 Run 'o list' to see all available shortcuts.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
