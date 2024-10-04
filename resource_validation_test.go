package main

import "testing"

func TestIsEmptyResource(t *testing.T) {
	// Test case 1: Both Memory and CPU are empty
	resource := compute{Memory: "", CPU: ""}
	if !isEmptyResource(resource) {
		t.Errorf("Expected resource to be empty")
	}

	// Test case 2: Memory is non-empty, CPU is empty
	resource = compute{Memory: "128Mi", CPU: ""}
	if isEmptyResource(resource) {
		t.Errorf("Expected resource not to be empty")
	}

	// Test case 3: Both Memory and CPU are non-empty
	resource = compute{Memory: "128Mi", CPU: "500m"}
	if isEmptyResource(resource) {
		t.Errorf("Expected resource not to be empty")
	}
}
