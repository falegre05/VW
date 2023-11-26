package tests

import (
	"testing"
	"vw/domain/grid"
	"vw/domain/robot"
)

// TestCreateGrid tests the CreateGrid function.
func TestCreateGrid(t *testing.T) {
	// Test case 1: Valid input
	input1 := "5 5"
	expectedGrid1 := &grid.Grid{XMax: 5, YMax: 5, Robots: make(map[string]*robot.Robot)}
	gridResult1, err1 := grid.CreateGrid(input1)

	// Check for errors
	if err1 != nil {
		t.Fatalf("Test case 1 failed. Unexpected error: %v", err1)
	}

	// Compare the actual result with the expected result
	if !isEqualGrid(gridResult1, expectedGrid1) {
		t.Errorf("Test case 1 failed. Got %+v, expected %+v", gridResult1, expectedGrid1)
	}

	// Test case 2: Invalid input
	input2 := "invalid input"
	_, err2 := grid.CreateGrid(input2)

	// Check for expected error
	if err2 == nil {
		t.Error("Test case 2 failed. Expected an error, but got none.")
	}
}

// Helper function to check equality of Grid instances
func isEqualGrid(got *grid.Grid, expected *grid.Grid) bool {
	// Check if XMax and YMax are equal
	if got.XMax != expected.XMax || got.YMax != expected.YMax {
		return false
	}

	// Check if the map lengths are equal
	if len(got.Robots) != len(expected.Robots) {
		return false
	}

	// Check each robot's position in the map
	for key, gotRobot := range got.Robots {
		expectedRobot, ok := expected.Robots[key]
		if !ok || !isEqualRobot(gotRobot, expectedRobot) {
			return false
		}
	}

	return true
}

// Helper function to check equality of Robot instances
func isEqualRobot(got *robot.Robot, expected *robot.Robot) bool {
	// Check if ID, X, Y, and Orientation are equal
	return got.ID == expected.ID &&
		got.X == expected.X &&
		got.Y == expected.Y &&
		got.Orientation == expected.Orientation
}
