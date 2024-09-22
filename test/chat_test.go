package test

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestIntegration_Chat(t *testing.T) {
	cmd := exec.Command("go", "run", "../cmd/chat/main.go")
	cmd.Env = append(
		cmd.Environ(),
		"LLM_PROMPT=What are the monthly sales for 2013?",
		"ENV_FILE_PATH=../internal/config/.env",
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to run chat: %v\nstdout: %s\nstderr: %s", err, stdout.String(), stderr.String())
	}

	t.Logf("stdout: %s", stdout.String())
	t.Logf("stderr: %s", stderr.String())
}
