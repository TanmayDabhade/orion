package cmd

import (
	"context"
	"strings"

	"orion/internal/ai"
	"orion/internal/config"
	"orion/internal/history"

	"github.com/spf13/cobra"
)

var aiCmd = &cobra.Command{
	Use:   "ai <task>",
	Short: "Route a task to the AI provider",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := strings.Join(args, " ")
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		store, err := history.Open(config.HistoryPath())
		if err != nil {
			return err
		}
		defer store.Close()

		intent, err := ai.InferIntent(context.Background(), input, cfg)
		if err != nil {
			return err
		}

		return executeIntent(cmd, store, input, intent, cfg)
	},
}

func init() {
	rootCmd.AddCommand(aiCmd)
}
