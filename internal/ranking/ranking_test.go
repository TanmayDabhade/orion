package ranking

import (
	"testing"

	"orion/internal/history"
)

func TestRankedKeys(t *testing.T) {
	keys := []string{"alpha", "beta", "gamma"}
	usage := map[string]history.Usage{
		"beta":  {Key: "beta", Count: 3, LastUnix: 10},
		"alpha": {Key: "alpha", Count: 2, LastUnix: 20},
	}

	ranked := RankedKeys(keys, usage)
	if ranked[0] != "beta" {
		t.Fatalf("expected beta first, got %s", ranked[0])
	}
}
