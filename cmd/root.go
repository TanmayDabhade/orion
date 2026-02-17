package cmd

import (
	"context"
	"os"
	"strings"

	"orion/internal/ai"
	"orion/internal/config"
	"orion/internal/executor"
	"orion/internal/history"
	"orion/internal/plan"
	"orion/internal/router"
	"orion/internal/safety"
	"orion/internal/shortcuts"
	"orion/models"

	"github.com/spf13/cobra"
)

var (
	yes     bool
	Version = "dev" // overridden by ldflags
)

var rootCmd = &cobra.Command{
	Use:     "o",
	Version: Version,
	Short:   "Orion: natural-language terminal assistant",
	Args:    cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}

		input := strings.Join(args, " ")
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		// Ensure config directory exists for DB
		if err := os.MkdirAll(config.Dir(), 0755); err != nil {
			return err
		}

		store, err := history.Open(config.HistoryPath())
		if err != nil {
			return err
		}
		defer store.Close()

		shortcutMap, err := shortcuts.Load(config.ShortcutsPath())
		if err != nil {
			return err
		}

		route, err := router.Route(input, shortcutMap)
		if err != nil {
			return err
		}

		intent := route.Intent
		if route.FallbackSearch && cfg.Features["ai_fallback"] {
			aiIntent, err := ai.InferIntent(context.Background(), input, cfg)
			if err != nil {
				return err
			}
			intent = aiIntent
		}

		return executeIntent(cmd, store, input, intent, cfg)
	},
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&yes, "yes", "y", false, "skip confirmations")
}

func executeIntent(cmd *cobra.Command, store *history.Store, input string, intent models.Intent, cfg config.Config) error {
	localRisk := safety.Assess(intent)
	if intent.Risk == "" {
		intent.Risk = localRisk
	} else if models.RiskRank(intent.Risk) < models.RiskRank(localRisk) {
		intent.Risk = localRisk
	}

	threshold := models.ParseRisk(cfg.RiskThreshold)
	if err := safety.Gate(intent, threshold, yes); err != nil {
		return err
	}

	commandPlan := plan.FromIntent(intent, cfg)
	err := executor.Execute(commandPlan, cfg)
	recordErr := store.Record(input, shortcuts.Normalize(input), err == nil)

	if err != nil {
		return err
	}
	if recordErr != nil {
		return recordErr
	}
	return nil
}
