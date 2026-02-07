package cmd

import (
	"fmt"
	"sort"

	"orion/internal/apps"
	"orion/internal/config"
	"orion/internal/history"
	"orion/internal/ranking"
	"orion/internal/shortcuts"

	"github.com/spf13/cobra"
)

var (
	ranked      bool
	listApps    bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List shortcuts or apps",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if listApps {
			fmt.Println("Scanning for applications...")
			appList := apps.List()
			if len(appList) == 0 {
				fmt.Println("No applications found.")
				return nil
			}

			// Sort by key
			keys := make([]string, 0, len(appList))
			for k := range appList {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for _, k := range keys {
				// Format: googlechrome -> /Applications/Google Chrome.app
				fmt.Printf("%s -> %s\n", k, appList[k])
			}
			return nil
		}

		entries, err := shortcuts.Load(config.ShortcutsPath())
		if err != nil {
			return err
		}

		keys := shortcuts.SortedKeys(entries)
		if len(keys) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No shortcuts found.")
			return nil
		}

		if ranked {
			store, err := history.Open(config.HistoryPath())
			if err != nil {
				return err
			}
			defer store.Close()

			usage, err := store.Usage(normalizedKeys(keys))
			if err != nil {
				return err
			}
			keys = ranking.RankedKeys(keys, usage)
		}

		for _, key := range keys {
			fmt.Fprintf(cmd.OutOrStdout(), "%s -> %s\n", key, entries[key])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&ranked, "ranked", false, "rank shortcuts by usage")
	listCmd.Flags().BoolVar(&listApps, "apps", false, "list detected applications")
}

func normalizedKeys(keys []string) []string {
	normalized := make([]string, 0, len(keys))
	for _, key := range keys {
		normalized = append(normalized, shortcuts.Normalize(key))
	}
	return normalized
}
