package tests

import (
	"testing"
	"vw/domain/grid"
	"vw/domain/instructions"
	"vw/domain/robot"
	"vw/domain/service"
)

// TestApplyInstructions tests the ApplyInstructions function of RobotService.
func TestApplyInstructions(t *testing.T) {
	// Create a new grid
	g, err := grid.CreateGrid("5 5")
	if err != nil {
		t.Fatalf("Failed to create grid: %v", err)
	}

	// Create a new RobotService with the grid
	rs := service.NewRobotService(g)

	// Test case 1: Apply instructions to move robot within grid limits
	robot1 := &robot.Robot{ID: 1, Orientation: robot.NorthOrientation, Position: robot.Position{X: 1, Y: 1}}
	instructions1 := instructions.InstructionSet{instructions.MoveForward, instructions.RotateLeft, instructions.MoveForward}
	err1 := rs.ApplyInstructions(robot1, instructions1)

	// Check for errors
	if err1 != nil {
		t.Errorf("Test case 1 failed. Unexpected error: %v", err1)
	}

	// Check the final position of the robot
	expectedPosition1 := &robot.Position{X: 0, Y: 2}
	if !robotPositionsEqual(&robot1.Position, expectedPosition1) {
		t.Errorf("Test case 1 failed. Got final position %+v, expected %+v", robot1.Position, expectedPosition1)
	}

	// Test case 2: Apply instructions to move robot outside grid limits
	robot2 := &robot.Robot{ID: 2, Orientation: robot.EastOrientation, Position: robot.Position{X: 5, Y: 5}}
	instructions2 := instructions.InstructionSet{instructions.MoveForward, instructions.MoveForward}
	err2 := rs.ApplyInstructions(robot2, instructions2)

	// Check for expected error
	if err2 == nil {
		t.Error("Test case 2 failed. Expected an error, but got none.")
	}

	// Test case 3: Apply instructions to move robot to an occupied position
	robot3 := &robot.Robot{ID: 3, Orientation: robot.SouthOrientation, Position: robot.Position{X: 3, Y: 3}}
	// Add another robot to the same position
	g.Robots[robot3.Position.String()] = &robot.Robot{ID: 4}
	instructions3 := instructions.InstructionSet{instructions.MoveForward}
	err3 := rs.ApplyInstructions(robot3, instructions3)

	// Check for expected error
	if err3 == nil {
		t.Error("Test case 3 failed. Expected an error, but got none.")
	}
}

// robotPositionsEqual checks if two robot positions are equal.
func robotPositionsEqual(pos1, pos2 *robot.Position) bool {
	return pos1.X == pos2.X && pos1.Y == pos2.Y
}
