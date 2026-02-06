package models

type Action string

const (
	ActionOpenURL  Action = "open_url"
	ActionOpenApp  Action = "open_app"
	ActionSearch   Action = "search"
	ActionRunShell Action = "run_shell"
	ActionFileFind Action = "file_find"
	ActionGitHelp  Action = "git_helper"
	ActionAI       Action = "ai"
)

type Intent struct {
	Action Action
	Args   map[string]string
	Risk   Risk
}
