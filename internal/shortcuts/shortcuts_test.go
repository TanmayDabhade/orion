package shortcuts

import "testing"

func TestResolve(t *testing.T) {
	entries := map[string]string{
		"msu d2l": "open https://d2l.msu.edu",
	}

	if cmd, ok := Resolve(entries, "MSU    D2L"); !ok || cmd == "" {
		t.Fatalf("expected shortcut to resolve")
	}
}
