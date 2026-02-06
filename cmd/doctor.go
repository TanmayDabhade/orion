package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"orion/internal/ai"
	"orion/internal/config"
	"orion/internal/history"
	"orion/internal/shortcuts"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Orion environment",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		checks := []struct {
			name string
			fn   func() error
		}{
			{
				name: "Config directory",
				fn: func() error {
					return os.MkdirAll(config.Dir(), 0o755)
				},
			},
			{
				name: "Config file",
				fn: func() error {
					_, err := os.Stat(config.Path())
					if os.IsNotExist(err) {
						return nil
					}
					return err
				},
			},
			{
				name: "Shortcuts file",
				fn: func() error {
					_, err := shortcuts.Load(config.ShortcutsPath())
					return err
				},
			},
			{
				name: "Search engine",
				fn: func() error {
					if !strings.Contains(cfg.SearchEngine, "%s") {
						return fmt.Errorf("search_engine must include %%s placeholder")
					}
					return nil
				},
			},
			{
				name: "History database",
				fn: func() error {
					store, err := history.Open(config.HistoryPath())
					if err != nil {
						return err
					}
					return store.Close()
				},
			},
			{
				name: "AI provider",
				fn: func() error {
					return ai.HealthCheck(context.Background(), cfg)
				},
			},
		}

		failures := 0
		for _, check := range checks {
			if err := check.fn(); err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "[FAIL] %s: %v\n", check.name, err)
				failures++
				continue
			}
			fmt.Fprintf(cmd.OutOrStdout(), "[OK] %s\n", check.name)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Config: %s\n", config.Path())
		fmt.Fprintf(cmd.OutOrStdout(), "Shortcuts: %s\n", config.ShortcutsPath())
		fmt.Fprintf(cmd.OutOrStdout(), "AI Provider: %s\n", cfg.AIProvider)

		if failures > 0 {
			return fmt.Errorf("doctor found %d issue(s)", failures)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
