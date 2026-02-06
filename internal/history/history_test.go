package history

import (
	"path/filepath"
	"testing"

	"orion/internal/shortcuts"
)

func TestRecordAndUsage(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "history.db")
	store, err := Open(path)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer store.Close()

	input := "MSU   D2L"
	key := shortcuts.Normalize(input)
	if err := store.Record(input, key, true); err != nil {
		t.Fatalf("record: %v", err)
	}

	usage, err := store.Usage([]string{key})
	if err != nil {
		t.Fatalf("usage: %v", err)
	}
	if usage[key].Count != 1 {
		t.Fatalf("expected count 1, got %d", usage[key].Count)
	}
}
