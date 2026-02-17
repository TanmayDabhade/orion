package local

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLocal_Complete(t *testing.T) {
	// Create a mock sidecar script that just echoes a JSON plan
	tmpDir := t.TempDir()
	mockBin := filepath.Join(tmpDir, "mock-llm.sh")

	script := `#!/bin/sh
cat > /dev/null # consume stdin
echo '{"intent":"test","summary":"Mock plan","commands":[{"cmd":"echo success","risk":"low"}],"questions":[]}'
`
	if err := os.WriteFile(mockBin, []byte(script), 0755); err != nil {
		t.Fatalf("Failed to create mock binary: %v", err)
	}

	provider := Local{BinPath: mockBin}
	ctx := context.Background()

	output, err := provider.Complete(ctx, "some prompt")
	if err != nil {
		t.Fatalf("Complete failed: %v", err)
	}

	if !strings.Contains(output, "echo success") {
		t.Errorf("Expected output to contain 'echo success', got: %s", output)
	}
}

func TestLocal_Health(t *testing.T) {
	tmpDir := t.TempDir()
	mockBin := filepath.Join(tmpDir, "exists.sh")
	if err := os.WriteFile(mockBin, []byte("#!/bin/sh"), 0755); err != nil {
		t.Fatal(err)
	}

	provider := Local{BinPath: mockBin}
	if err := provider.Health(context.Background()); err != nil {
		t.Errorf("Health check failed for existing binary: %v", err)
	}

	badProvider := Local{BinPath: "/does/not/exist"}
	if err := badProvider.Health(context.Background()); err == nil {
		t.Error("Health check should have failed for missing binary")
	}
}
