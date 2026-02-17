package plan

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ParseStrict parses a JSON byte slice into a CommandPlan with strict rules:
// - No unknown fields
// - Max 8 commands
// - Valid risk levels
func ParseStrict(input []byte) (*CommandPlan, error) {
	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.DisallowUnknownFields()

	var plan CommandPlan
	if err := decoder.Decode(&plan); err != nil {
		return nil, fmt.Errorf("invalid plan JSON: %w", err)
	}

	if err := validatePlan(&plan); err != nil {
		return nil, err
	}

	return &plan, nil
}

func validatePlan(plan *CommandPlan) error {
	if plan.Intent == "" {
		return errors.New("missing intent")
	}
	// A plan must do something: either run commands or ask clarifying questions
	if len(plan.Commands) == 0 && len(plan.Questions) == 0 {
		return errors.New("plan must have commands or questions")
	}
	if len(plan.Commands) > 8 {
		return errors.New("too many commands (max 8)")
	}

	for i, cmd := range plan.Commands {
		if strings.TrimSpace(cmd.Cmd) == "" {
			return fmt.Errorf("command %d is empty", i)
		}

		switch cmd.Risk {
		case RiskLow, RiskMedium, RiskHigh:
		default:
			return fmt.Errorf("command %d has invalid risk: %s", i, cmd.Risk)
		}
	}

	return nil
}
