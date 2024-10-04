package main

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestValidateImmutable_TableDriven(t *testing.T) {
	tests := []struct {
		name        string
		pod         podSpec
		expectedLog string
	}{
		{
			name: "ConfigMap",
			pod: podSpec{
				Kind:      "ConfigMap",
				Metadata:  metadata{Name: "test-config-map"},
				Immutable: false,
			},
			expectedLog: "ConfigMap test-config-map should be immutable.",
		},
		{
			name: "Secret",
			pod: podSpec{
				Kind:      "Secret",
				Metadata:  metadata{Name: "test-secret"},
				Immutable: false,
			},
			expectedLog: "Secret test-secret should be immutable.",
		},
		{
			name: "ImmutablePod",
			pod: podSpec{
				Kind:      "ConfigMap",
				Metadata:  metadata{Name: "test-config-map"},
				Immutable: true,
			},
			expectedLog: "",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			var logOutput bytes.Buffer
			log.SetOutput(&logOutput)

			// Call the function under test
			validateImmutable(testcase.pod)

			// Check log output
			if !strings.Contains(logOutput.String(), testcase.expectedLog) {
				t.Errorf("Expected log: %q, but got: %q", testcase.expectedLog, logOutput.String())
			}
		})
	}

	// Restore the default log output after all tests
	log.SetOutput(nil)
}
