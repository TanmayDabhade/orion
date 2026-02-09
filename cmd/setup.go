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

		// Merge: app shortcuts take lower priority than user-defined ones
		for key, path := range appMap {
			if _, exists := existingShortcuts[key]; !exists {
				existingShortcuts[key] = fmt.Sprintf("open -a '%s'", path)
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
