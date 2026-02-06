package models

import "strings"

type Risk string

const (
	RiskLow    Risk = "low"
	RiskMedium Risk = "medium"
	RiskHigh   Risk = "high"
)

func ParseRisk(s string) Risk {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case string(RiskLow):
		return RiskLow
	case string(RiskHigh):
		return RiskHigh
	default:
		return RiskMedium
	}
}

func RiskRank(r Risk) int {
	switch r {
	case RiskLow:
		return 1
	case RiskHigh:
		return 3
	default:
		return 2
	}
}
