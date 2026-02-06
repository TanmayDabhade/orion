package safety

import (
	"testing"

	"orion/models"
)

func TestAssess(t *testing.T) {
	high := models.Intent{Action: models.ActionRunShell, Args: map[string]string{"command": "rm -rf /"}}
	if risk := Assess(high); risk != models.RiskHigh {
		t.Fatalf("expected high risk, got %s", risk)
	}

	medium := models.Intent{Action: models.ActionRunShell, Args: map[string]string{"command": "rm -rf ./tmp"}}
	if risk := Assess(medium); risk != models.RiskMedium {
		t.Fatalf("expected medium risk, got %s", risk)
	}

	low := models.Intent{Action: models.ActionOpenURL, Args: map[string]string{"url": "https://example.com"}}
	if risk := Assess(low); risk != models.RiskLow {
		t.Fatalf("expected low risk, got %s", risk)
	}
}
