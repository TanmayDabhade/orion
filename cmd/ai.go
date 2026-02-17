package cmd

import (
	"context"
	"fmt"
	"strings"

	"orion/internal/ai"
	"orion/internal/config"
	"orion/internal/executor"
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

		// Interactive setup for new users
		if !cfg.Features["ai_setup_done"] {
			// Only prompt if no key is present
			if cfg.AIKey == "" {
				fmt.Println("⚠️  No AI provider configured.")
				fmt.Print("Enter Gemini API Key (or press Enter to use local Ollama): ")

				var key string
				if _, err := fmt.Scanln(&key); err != nil && err.Error() != "unexpected newline" {
					// ignore
				}

				key = strings.TrimSpace(key)
				if key != "" {
					cfg.AIKey = key
					cfg.AIProvider = "gemini"
					fmt.Println("✅ Configured Gemini.")
				} else {
					cfg.AIProvider = "ollama"
					fmt.Println("ℹ️  Using local Ollama.")
				}
			}

			// Mark setup as done
			if cfg.Features == nil {
				cfg.Features = make(map[string]bool)
			}
			cfg.Features["ai_setup_done"] = true
			if err := config.Save(cfg); err != nil {
				fmt.Printf("Warning: failed to save config: %v\n", err)
			}
		}

		store, err := history.Open(config.HistoryPath())
		if err != nil {
			return err
		}
		defer store.Close()

		fmt.Println("🤔 Thinking...")
		plan, err := ai.InferPlan(context.Background(), input, cfg)
		if err != nil {
			return err
		}

		// Confirm execution
		fmt.Printf("\n📋 Plan: %s\n", plan.Summary)
		for _, q := range plan.Questions {
			fmt.Printf("❓ Question: %s\n", q)
		}
		if len(plan.Commands) > 0 {
			fmt.Println("Commands to run:")
			for i, cmd := range plan.Commands {
				fmt.Printf("  %d. %s (%s)\n", i+1, cmd.Cmd, cmd.Risk)
			}

			if !config.Confirm("Execute this plan?") {
				return fmt.Errorf("aborted by user")
			}

			fmt.Println("\n🚀 Executing...")
			if err := executor.Execute(plan, cfg); err != nil {
				return err
			}
		} else {
			fmt.Println("No commands to execute.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(aiCmd)
}
