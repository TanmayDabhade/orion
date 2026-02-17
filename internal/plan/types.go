package plan

type RiskLevel string

const (
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium"
	RiskHigh   RiskLevel = "high"
)

type Command struct {
	Cmd  string    `json:"cmd"`
	Risk RiskLevel `json:"risk"`
}

type CommandPlan struct {
	Intent    string    `json:"intent"`
	Summary   string    `json:"summary"`
	Cwd       string    `json:"cwd"`
	Commands  []Command `json:"commands"`
	Questions []string  `json:"questions"`
}
