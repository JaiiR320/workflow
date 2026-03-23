package main

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestStoreCLI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	binaryPath := buildStoreBinary(t)
	workingDir := t.TempDir()
	key := "key with spaces"
	value := "value with \"quotes\" & symbols !@#$%^&*()"

	stdout, stderr, exitCode := runStore(t, binaryPath, workingDir, "set", key, value)
	if exitCode != 0 {
		t.Fatalf("store set exit code = %d, want 0; stderr = %q", exitCode, stderr)
	}
	if got, want := strings.TrimSpace(stdout), "stored "+key; got != want {
		t.Fatalf("store set stdout = %q, want %q", got, want)
	}
	if stderr != "" {
		t.Fatalf("store set stderr = %q, want empty", stderr)
	}

	stored := readStoreFile(t, filepath.Join(workingDir, storeFileName))
	if got := stored[key]; got != value {
		t.Fatalf("stored value for %q = %q, want %q", key, got, value)
	}

	stdout, stderr, exitCode = runStore(t, binaryPath, workingDir, "get", key)
	if exitCode != 0 {
		t.Fatalf("store get exit code = %d, want 0; stderr = %q", exitCode, stderr)
	}
	if got, want := strings.TrimSpace(stdout), value; got != want {
		t.Fatalf("store get stdout = %q, want %q", got, want)
	}
	if stderr != "" {
		t.Fatalf("store get stderr = %q, want empty", stderr)
	}

	stdout, stderr, exitCode = runStore(t, binaryPath, workingDir, "list")
	if exitCode != 0 {
		t.Fatalf("store list exit code = %d, want 0; stderr = %q", exitCode, stderr)
	}
	if !strings.Contains(stdout, "KEY") || !strings.Contains(stdout, value) || !strings.Contains(stdout, key) {
		t.Fatalf("store list stdout = %q, want header and stored entry", stdout)
	}
	if stderr != "" {
		t.Fatalf("store list stderr = %q, want empty", stderr)
	}

	stdout, stderr, exitCode = runStore(t, binaryPath, workingDir, "delete", key)
	if exitCode != 0 {
		t.Fatalf("store delete exit code = %d, want 0; stderr = %q", exitCode, stderr)
	}
	if got, want := strings.TrimSpace(stdout), "deleted "+key; got != want {
		t.Fatalf("store delete stdout = %q, want %q", got, want)
	}
	if stderr != "" {
		t.Fatalf("store delete stderr = %q, want empty", stderr)
	}

	stdout, stderr, exitCode = runStore(t, binaryPath, workingDir, "get", key)
	if exitCode != 1 {
		t.Fatalf("store get missing exit code = %d, want 1", exitCode)
	}
	if stdout != "" {
		t.Fatalf("store get missing stdout = %q, want empty", stdout)
	}
	if got, want := strings.TrimSpace(stderr), "not found"; got != want {
		t.Fatalf("store get missing stderr = %q, want %q", got, want)
	}

	stdout, stderr, exitCode = runStore(t, binaryPath, workingDir, "delete", key)
	if exitCode != 1 {
		t.Fatalf("store delete missing exit code = %d, want 1", exitCode)
	}
	if stdout != "" {
		t.Fatalf("store delete missing stdout = %q, want empty", stdout)
	}
	if got, want := strings.TrimSpace(stderr), "not found"; got != want {
		t.Fatalf("store delete missing stderr = %q, want %q", got, want)
	}

	stored = readStoreFile(t, filepath.Join(workingDir, storeFileName))
	if len(stored) != 0 {
		t.Fatalf("remaining store entries = %v, want empty", stored)
	}
}

func TestStoreCLI_RejectsEmptyKey(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	binaryPath := buildStoreBinary(t)
	workingDir := t.TempDir()
	tests := []struct {
		name string
		args []string
	}{
		{name: "set", args: []string{"set", "", "value"}},
		{name: "get", args: []string{"get", ""}},
		{name: "delete", args: []string{"delete", ""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, exitCode := runStore(t, binaryPath, workingDir, tt.args...)
			if exitCode != 1 {
				t.Fatalf("%s exit code = %d, want 1", tt.name, exitCode)
			}
			if stdout != "" {
				t.Fatalf("%s stdout = %q, want empty", tt.name, stdout)
			}
			if got, want := strings.TrimSpace(stderr), "key cannot be empty"; got != want {
				t.Fatalf("%s stderr = %q, want %q", tt.name, got, want)
			}
		})
	}
}

func buildStoreBinary(t *testing.T) string {
	t.Helper()

	binaryPath := filepath.Join(t.TempDir(), "store-test")
	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	cmd.Dir = moduleRoot(t)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go build failed: %v\n%s", err, output)
	}

	return binaryPath
}

func moduleRoot(t *testing.T) string {
	t.Helper()

	root, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd failed: %v", err)
	}

	return root
}

func runStore(t *testing.T, binaryPath, workingDir string, args ...string) (string, string, int) {
	t.Helper()

	cmd := exec.Command(binaryPath, args...)
	cmd.Dir = workingDir
	stdout, err := cmd.Output()
	if err == nil {
		return string(stdout), "", 0
	}

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return string(stdout), string(exitErr.Stderr), exitErr.ExitCode()
	}

	t.Fatalf("running %q failed: %v", strings.Join(args, " "), err)
	return "", "", 0
}

func readStoreFile(t *testing.T, path string) storeData {
	t.Helper()

	contents, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("os.ReadFile(%q) failed: %v", path, err)
	}

	var data storeData
	if err := json.Unmarshal(contents, &data); err != nil {
		t.Fatalf("json.Unmarshal(%q) failed: %v", path, err)
	}

	return data
}
