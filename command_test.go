package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestValidateCommand(t *testing.T) {
	// Step 1: Create a buffer to capture the output
	var outputBuffer bytes.Buffer

	// Step 2: Set the output for the command to the buffer
	validateCmd.SetOut(&outputBuffer)

	// Step 3: Simulate passing arguments to the Cobra command
	validateCmd.SetArgs([]string{"--file=example.yaml"})

	// Step 4: Execute the command
	err := validateCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	// Step 5: Check the output
	output := outputBuffer.String()
	if !strings.Contains(output, "example.yaml") {
		t.Errorf("Expected output to contain 'example.yaml', but got: %v", output)
	}
}

func TestValidateCommand_MissingFileFlag(t *testing.T) {
	// Step 1: Create a buffer to capture the output
	var outputBuffer bytes.Buffer

	// Step 2: Set the output for the command to the buffer
	validateCmd.SetOut(&outputBuffer)

	// Step 3: Simulate passing invalid arguments (no --file flag)
	validateCmd.SetArgs([]string{})

	// Step 4: Execute the command
	err := validateCmd.Execute()

	// Step 5: Check if the command failed as expected
	if err == nil {
		t.Fatalf("Expected an error for missing --file flag, but got none")
	}

	// Step 6: Check the output for the error message
	output := outputBuffer.String()
	if !strings.Contains(output, "Error: --file flag is required.") {
		t.Errorf("Expected error message 'Error: --file flag is required.', but got: %v", output)
	}
}
