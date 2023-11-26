package tests

import (
	"reflect"
	"testing"
	"vw/domain/robot"
)

// TestCreateRobot tests the CreateRobot function.
func TestCreateRobot(t *testing.T) {
	// Test case 1: Valid input
	input1 := "1 2 N"
	expectedRobot1 := &robot.Robot{ID: 1, Orientation: robot.NorthOrientation, Position: robot.Position{X: 1, Y: 2}}
	robotResult1, err1 := robot.CreateRobot(1, input1)

	// Check for errors
	if err1 != nil {
		t.Fatalf("Test case 1 failed. Unexpected error: %v", err1)
	}

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(robotResult1, expectedRobot1) {
		t.Errorf("Test case 1 failed. Got %+v, expected %+v", robotResult1, expectedRobot1)
	}

	// Test case 2: Invalid input
	input2 := "invalid input"
	_, err2 := robot.CreateRobot(2, input2)

	// Check for expected error
	if err2 == nil {
		t.Error("Test case 2 failed. Expected an error, but got none.")
	}
}

// TestMove tests the Move function.
func TestMove(t *testing.T) {
	// Test case 1: Move north
	robot1 := &robot.Robot{Orientation: robot.NorthOrientation, Position: robot.Position{X: 1, Y: 2}}
	expectedPosition1 := robot.Position{X: 1, Y: 3}
	robot1.Move()

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(robot1.Position, expectedPosition1) {
		t.Errorf("Test case 1 failed. Got %+v, expected %+v", robot1.Position, expectedPosition1)
	}

	// Test case 2: Move east
	robot2 := &robot.Robot{Orientation: robot.EastOrientation, Position: robot.Position{X: 1, Y: 2}}
	expectedPosition2 := robot.Position{X: 2, Y: 2}
	robot2.Move()

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(robot2.Position, expectedPosition2) {
		t.Errorf("Test case 2 failed. Got %+v, expected %+v", robot2.Position, expectedPosition2)
	}

	// Test case 3: Move south
	robot3 := &robot.Robot{Orientation: robot.SouthOrientation, Position: robot.Position{X: 1, Y: 2}}
	expectedPosition3 := robot.Position{X: 1, Y: 1}
	robot3.Move()

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(robot3.Position, expectedPosition3) {
		t.Errorf("Test case 3 failed. Got %+v, expected %+v", robot3.Position, expectedPosition3)
	}

	// Test case 4: Move west
	robot4 := &robot.Robot{Orientation: robot.WestOrientation, Position: robot.Position{X: 1, Y: 2}}
	expectedPosition4 := robot.Position{X: 0, Y: 2}
	robot4.Move()

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(robot4.Position, expectedPosition4) {
		t.Errorf("Test case 4 failed. Got %+v, expected %+v", robot4.Position, expectedPosition4)
	}
}

// TestCalcPositionAfterMovingForward tests the CalcPositionAfterMovingForward function.
func TestCalcPositionAfterMovingForward(t *testing.T) {
	// Test case 1: Move forward facing north
	robot1 := &robot.Robot{Orientation: robot.NorthOrientation, Position: robot.Position{X: 1, Y: 2}}
	expectedPosition1 := &robot.Position{X: 1, Y: 3}
	result1 := robot1.CalcPositionAfterMovingForward()

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(result1, expectedPosition1) {
		t.Errorf("Test case 1 failed. Got %+v, expected %+v", result1, expectedPosition1)
	}

	// Test case 2: Move forward facing east
	robot2 := &robot.Robot{Orientation: robot.EastOrientation, Position: robot.Position{X: 1, Y: 2}}
	expectedPosition2 := &robot.Position{X: 2, Y: 2}
	result2 := robot2.CalcPositionAfterMovingForward()

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(result2, expectedPosition2) {
		t.Errorf("Test case 2 failed. Got %+v, expected %+v", result2, expectedPosition2)
	}

	// Test case 3: Move forward facing south
	robot3 := &robot.Robot{Orientation: robot.SouthOrientation, Position: robot.Position{X: 1, Y: 2}}
	expectedPosition3 := &robot.Position{X: 1, Y: 1}
	result3 := robot3.CalcPositionAfterMovingForward()

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(result3, expectedPosition3) {
		t.Errorf("Test case 3 failed. Got %+v, expected %+v", result3, expectedPosition3)
	}

	// Test case 4: Move forward facing west
	robot4 := &robot.Robot{Orientation: robot.WestOrientation, Position: robot.Position{X: 1, Y: 2}}
	expectedPosition4 := &robot.Position{X: 0, Y: 2}
	result4 := robot4.CalcPositionAfterMovingForward()

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(result4, expectedPosition4) {
		t.Errorf("Test case 2 failed. Got %+v, expected %+v", result4, expectedPosition4)
	}
}

// TestTurnLeft tests the TurnLeft function.
func TestTurnLeft(t *testing.T) {
	// Test case 1: Turn left from north
	robot1 := &robot.Robot{Orientation: robot.NorthOrientation}
	expectedOrientation1 := robot.WestOrientation
	robot1.TurnLeft()

	// Compare the actual result with the expected result
	if robot1.Orientation != expectedOrientation1 {
		t.Errorf("Test case 1 failed. Got %s, expected %s", robot1.Orientation, expectedOrientation1)
	}

	// Test case 2: Turn left from east
	robot2 := &robot.Robot{Orientation: robot.EastOrientation}
	expectedOrientation2 := robot.NorthOrientation
	robot2.TurnLeft()

	// Compare the actual result with the expected result
	if robot2.Orientation != expectedOrientation2 {
		t.Errorf("Test case 2 failed. Got %s, expected %s", robot2.Orientation, expectedOrientation2)
	}

	// Test case 3: Turn left from south
	robot3 := &robot.Robot{Orientation: robot.SouthOrientation}
	expectedOrientation3 := robot.EastOrientation
	robot3.TurnLeft()

	// Compare the actual result with the expected result
	if robot3.Orientation != expectedOrientation3 {
		t.Errorf("Test case 3 failed. Got %s, expected %s", robot3.Orientation, expectedOrientation3)
	}

	// Test case 4: Turn left from west
	robot4 := &robot.Robot{Orientation: robot.WestOrientation}
	expectedOrientation4 := robot.SouthOrientation
	robot4.TurnLeft()

	// Compare the actual result with the expected result
	if robot4.Orientation != expectedOrientation4 {
		t.Errorf("Test case 4 failed. Got %s, expected %s", robot4.Orientation, expectedOrientation4)
	}
}

// TestTurnRight tests the TurnRight function.
func TestTurnRight(t *testing.T) {
	// Test case 1: Turn right from north
	robot1 := &robot.Robot{Orientation: robot.NorthOrientation}
	expectedOrientation1 := robot.EastOrientation
	robot1.TurnRight()

	// Compare the actual result with the expected result
	if robot1.Orientation != expectedOrientation1 {
		t.Errorf("Test case 1 failed. Got %s, expected %s", robot1.Orientation, expectedOrientation1)
	}

	// Test case 2: Turn right from west
	robot2 := &robot.Robot{Orientation: robot.WestOrientation}
	expectedOrientation2 := robot.NorthOrientation
	robot2.TurnRight()

	// Compare the actual result with the expected result
	if robot2.Orientation != expectedOrientation2 {
		t.Errorf("Test case 2 failed. Got %s, expected %s", robot2.Orientation, expectedOrientation2)
	}

	// Test case 3: Turn right from south
	robot3 := &robot.Robot{Orientation: robot.SouthOrientation}
	expectedOrientation3 := robot.WestOrientation
	robot3.TurnRight()

	// Compare the actual result with the expected result
	if robot3.Orientation != expectedOrientation3 {
		t.Errorf("Test case 1 failed. Got %s, expected %s", robot3.Orientation, expectedOrientation3)
	}

	// Test case 4: Turn right from west
	robot4 := &robot.Robot{Orientation: robot.WestOrientation}
	expectedOrientation4 := robot.NorthOrientation
	robot4.TurnRight()

	// Compare the actual result with the expected result
	if robot4.Orientation != expectedOrientation4 {
		t.Errorf("Test case 2 failed. Got %s, expected %s", robot4.Orientation, expectedOrientation4)
	}
}
