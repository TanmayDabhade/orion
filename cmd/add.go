package cmd

import (
	"fmt"

	"orion/internal/config"
	"orion/internal/shortcuts"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <phrase> <command>",
	Short: "Add a shortcut",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		phrase := args[0]
		command := args[1]
		for i := 2; i < len(args); i++ {
			command += " " + args[i]
		}

		path := config.ShortcutsPath()
		entries, err := shortcuts.Load(path)
		if err != nil {
			return err
		}

		entries[phrase] = command
		if err := shortcuts.Save(path, entries); err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Added shortcut: %s -> %s\n", phrase, command)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
