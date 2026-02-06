package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"orion/internal/config"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit shortcuts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := config.ShortcutsPath()
		if err := os.MkdirAll(config.Dir(), 0o755); err != nil {
			return err
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.WriteFile(path, []byte("shortcuts:\n"), 0o644); err != nil {
				return err
			}
		}

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = os.Getenv("VISUAL")
		}

		var cmdExec *exec.Cmd
		if editor != "" {
			cmdExec = exec.Command(editor, path)
		} else {
			cmdExec = exec.Command("open", "-e", path)
		}

		cmdExec.Stdout = cmd.OutOrStdout()
		cmdExec.Stderr = cmd.ErrOrStderr()
		cmdExec.Stdin = os.Stdin
		if err := cmdExec.Run(); err != nil {
			return fmt.Errorf("launch editor: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
