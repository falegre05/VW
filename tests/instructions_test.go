package tests

import (
	"reflect"
	"testing"
	"vw/domain/instructions"
)

// TestParseInstructions tests the ParseInstructions function.
func TestParseInstructions(t *testing.T) {
	// Test case 1: Valid input
	input1 := "RML"
	expectedInstructions1 := instructions.InstructionSet{instructions.RotateRight, instructions.MoveForward, instructions.RotateLeft}
	instructionsResult1, err1 := instructions.ParseInstructions(input1)

	// Check for errors
	if err1 != nil {
		t.Fatalf("Test case 1 failed. Unexpected error: %v", err1)
	}

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(instructionsResult1, expectedInstructions1) {
		t.Errorf("Test case 1 failed. Got %+v, expected %+v", instructionsResult1, expectedInstructions1)
	}

	// Test case 2: Invalid input
	input2 := "XYZ"
	_, err2 := instructions.ParseInstructions(input2)

	// Check for expected error
	if err2 == nil {
		t.Error("Test case 2 failed. Expected an error, but got none.")
	}
}
