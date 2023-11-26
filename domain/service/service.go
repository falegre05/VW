package service

import (
	"fmt"
	"vw/domain/grid"
	"vw/domain/instructions"
	"vw/domain/robot"
)

// RobotService provides services related to robots.
type RobotService struct {
	Grid *grid.Grid
}

// NewRobotService creates a new instance of RobotService.
func NewRobotService(g *grid.Grid) *RobotService {
	return &RobotService{Grid: g}
}

// ApplyInstructions sequentially applies instructions for a robot.
// It updates the robot's position and orientation based on the provided instructions.
// If an invalid position is encountered, an error is returned.
func (rs *RobotService) ApplyInstructions(r *robot.Robot, instructionList instructions.InstructionSet) error {
	// Check if the first position is valid
	if err := rs.assertPositionIsValid(r, &r.Position); err != nil {
		return err
	}

	for _, inst := range instructionList {
		switch inst {
		case instructions.RotateLeft:
			r.TurnLeft()
		case instructions.RotateRight:
			r.TurnRight()
		case instructions.MoveForward:
			// Calculate the new position after moving forward
			newPosition := r.CalcPositionAfterMovingForward()

			// Check if the new position is valid
			if err := rs.assertPositionIsValid(r, newPosition); err != nil {
				return err
			}

			// Update the robot's position
			r.Move()
		}
	}

	// Print the final position and orientation of the robot
	fmt.Printf("#%d: %d %d %s\n", r.ID, r.X, r.Y, r.Orientation)

	// Mark the current position as occupied by the robot
	rs.Grid.Robots[r.Position.String()] = r

	return nil
}

// assertPositionIsValid checks whether the provided position for a robot is valid.
// It ensures that the position is within the grid limits and not already occupied by another robot.
func (rs *RobotService) assertPositionIsValid(r *robot.Robot, pos *robot.Position) error {
	switch {
	// Check if the X or Y coordinate of the position is outside the grid limits
	case pos.X < 0 || pos.X > rs.Grid.XMax || pos.Y < 0 || pos.Y > rs.Grid.YMax:
		return fmt.Errorf("robot #%d tried to move outside the grid limits at position %s", r.ID, pos)
	// Check if the position is already occupied by another robot
	case rs.Grid.Robots[pos.String()] != nil:
		return fmt.Errorf("robot #%d tried to move to the position %s, which is occupied by robot #%d", r.ID, pos, rs.Grid.Robots[pos.String()].ID)
	default:
		return nil
	}
}
