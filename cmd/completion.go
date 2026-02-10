package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:
  $ source <(o completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ o completion bash > /etc/bash_completion.d/o
  # macOS:
  $ o completion bash > /usr/local/etc/bash_completion.d/o

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ o completion zsh > "${fpath[1]}/_o"

  # You will need to start a new shell for this setup to take effect.

Fish:
  $ o completion fish | source

  # To load completions for each session, execute once:
  $ o completion fish > ~/.config/fish/completions/o.fish

PowerShell:
  PS> o completion powershell | Out-String | Invoke-Expression

  # To load completions for each session, execute once:
  PS> o completion powershell > o.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	// We replace the default completion command with our own custom one
	// But since Cobra adds the default one automatically, we don't need to add it again
	// unless we disable global completion which is complex.
	// Instead, we can just *override* it if we add it last, or we can just hope
	// to provide good docs.
	// Actually, Cobra's default completion command name is "completion".
	// If we add a command named "completion", it overrides the default one.
	rootCmd.AddCommand(completionCmd)
}
