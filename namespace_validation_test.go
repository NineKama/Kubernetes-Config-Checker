package main

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestValidateNamespace_TableDriven(t *testing.T) {
	tests := []struct {
		name        string
		pod         podSpec
		expectedLog string
	}{
		{
			name: "Pod with a dummy namespace",
			pod: podSpec{
				Kind:     "Pod",
				Metadata: metadata{Name: "test-pod", Namespace: "new-namespace"},
			},
			expectedLog: "",
		},
		{
			name: "Pod no namespace",
			pod: podSpec{
				Kind:      "Pod",
				Metadata:  metadata{Name: "test-pod"},
				Immutable: false,
			},
			expectedLog: "Pod test-pod should use a separate namespace (not default).",
		},
		{
			name: "Pod default namespace",
			pod: podSpec{
				Kind:      "Pod",
				Metadata:  metadata{Name: "test-pod", Namespace: "default"},
				Immutable: false,
			},
			expectedLog: "Pod test-pod should use a separate namespace (not default).",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			var logOutput bytes.Buffer
			log.SetOutput(&logOutput)

			// Call the function under test
			validateNameSpace(testcase.pod)

			// Check log output
			if !strings.Contains(logOutput.String(), testcase.expectedLog) {
				t.Errorf("Expected log: %q, but got: %q", testcase.expectedLog, logOutput.String())
			}

			logOutput.Reset()
		})
	}

	// Restore the default log output after all tests
	log.SetOutput(nil)
}
