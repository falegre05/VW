package tests

import (
	"log"
	"os"
	"reflect"
	"testing"
	"vw/domain/grid"
	"vw/domain/instructions"
	"vw/domain/robot"
	"vw/fileio"
)

// TestReadInstructionsFromFile tests the ReadInstructionsFromFile function.
func TestReadInstructionsFromFile(t *testing.T) {
	// Provide a test file path
	testFilePath := "test_input.txt"

	// Create a temporary test file with sample data
	createTestFile(testFilePath)

	// Clean up the temporary test file at the end
	defer cleanupTestFile(testFilePath)

	// Expected results based on the sample data in the test file
	expectedGrid := &grid.Grid{XMax: 5, YMax: 5, Robots: make(map[string]*robot.Robot)}
	expectedRobots := []robot.Robot{
		{ID: 0, Orientation: "N", Position: robot.Position{X: 1, Y: 2}},
		{ID: 1, Orientation: "E", Position: robot.Position{X: 3, Y: 3}},
	}
	expectedInstructions := []instructions.InstructionSet{
		{instructions.RotateLeft, instructions.MoveForward, instructions.RotateLeft, instructions.MoveForward, instructions.RotateLeft, instructions.MoveForward, instructions.RotateLeft, instructions.MoveForward, instructions.MoveForward},
		{instructions.MoveForward, instructions.MoveForward, instructions.RotateRight, instructions.MoveForward, instructions.MoveForward, instructions.RotateRight, instructions.MoveForward, instructions.RotateRight, instructions.RotateRight, instructions.MoveForward},
	}

	// Invoke the function under test
	gridResult, robotsResult, instructionsResult, err := fileio.ReadInstructionsFromFile(testFilePath)

	// Check for errors
	if err != nil {
		t.Fatalf("ReadInstructionsFromFile failed with error: %v", err)
	}

	// Compare the actual results with the expected results
	if !reflect.DeepEqual(gridResult, expectedGrid) {
		t.Errorf("Grid mismatch. Got %+v, expected %+v", gridResult, expectedGrid)
	}

	if !reflect.DeepEqual(robotsResult, expectedRobots) {
		t.Errorf("Robots mismatch. Got %+v, expected %+v", robotsResult, expectedRobots)
	}

	if !reflect.DeepEqual(instructionsResult, expectedInstructions) {
		t.Errorf("Instructions mismatch. Got %+v, expected %+v", instructionsResult, expectedInstructions)
	}
}

// Helper function to create a temporary test file with sample data
func createTestFile(filePath string) {
	content := "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Helper function to clean up the temporary test file
func cleanupTestFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
