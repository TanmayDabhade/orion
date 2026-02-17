package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"orion/internal/config"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [last|tail]",
	Short: "View recent execution logs",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		logsDir := config.LogsDir()
		entries, err := os.ReadDir(logsDir)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("No logs found.")
				return nil
			}
			return err
		}

		var files []string
		for _, e := range entries {
			if !e.IsDir() && filepath.Ext(e.Name()) == ".log" {
				files = append(files, filepath.Join(logsDir, e.Name()))
			}
		}

		if len(files) == 0 {
			fmt.Println("No logs found.")
			return nil
		}

		// Sort by modification time (newest last)
		// Since we name files by date, sorting by name is sufficient
		sort.Strings(files)

		if len(args) > 0 && args[0] == "last" {
			lastFile := files[len(files)-1]
			fmt.Printf("Displaying log: %s\n\n", lastFile)
			return catFile(lastFile)
		}

		fmt.Println("Recent logs:")
		for _, f := range files {
			fmt.Printf("  %s\n", filepath.Base(f))
		}
		fmt.Println("\nUse 'o logs last' to see the latest log.")
		return nil
	},
}

func catFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(os.Stdout, f)
	return err
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
