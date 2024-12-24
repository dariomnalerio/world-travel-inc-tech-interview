package main

import (
	"testing"
)

// sample test for setting up github actions
func TestSample(t *testing.T) {
	expected := true
	actual := true

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
