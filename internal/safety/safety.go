package safety

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"orion/models"
)

func Assess(intent models.Intent) models.Risk {
	switch intent.Action {
	case models.ActionRunShell:
		cmd := strings.ToLower(intent.Args["command"])
		if isHighRiskCommand(cmd) {
			return models.RiskHigh
		}
		if isMediumRiskCommand(cmd) {
			return models.RiskMedium
		}
		return models.RiskLow
	case models.ActionOpenURL, models.ActionOpenApp, models.ActionSearch:
		return models.RiskLow
	default:
		return models.RiskLow
	}
}

func Gate(intent models.Intent, threshold models.Risk, autoYes bool) error {
	risk := intent.Risk
	if risk == "" {
		risk = Assess(intent)
	}

	if risk == models.RiskHigh {
		if autoYes {
			return nil
		}
		return fmt.Errorf("high-risk command blocked; rerun with --yes to override")
	}

	if models.RiskRank(risk) < models.RiskRank(threshold) || autoYes {
		return nil
	}

	ok, err := confirm(fmt.Sprintf("Proceed with %s action?", intent.Action))
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("aborted")
	}
	return nil
}

func confirm(prompt string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "%s [y/N]: ", prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	text = strings.ToLower(strings.TrimSpace(text))
	return text == "y" || text == "yes", nil
}

func isHighRiskCommand(cmd string) bool {
	highPatterns := []string{
		"rm -rf /",
		"rm -fr /",
		"--no-preserve-root",
		"mkfs",
		"dd if=",
		"diskutil erase",
	}
	for _, pattern := range highPatterns {
		if strings.Contains(cmd, pattern) {
			return true
		}
	}
	return false
}

func isMediumRiskCommand(cmd string) bool {
	mediumTokens := []string{
		"rm ",
		"sudo ",
		"chmod ",
		"chown ",
		"kill ",
		"mv ",
		"cp ",
	}
	for _, token := range mediumTokens {
		if strings.Contains(cmd, " "+token) || strings.HasPrefix(cmd, token) {
			return true
		}
	}
	return false
}
