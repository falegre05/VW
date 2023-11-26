package robot

import (
	"fmt"
	"strconv"
	"strings"
)

// Robot represents the domain model for a cleaning robot.
type Robot struct {
	ID int
	Orientation
	Position
}

// Orientation defines the possible orientations of a robot.
type Orientation string

const (
	NorthOrientation Orientation = "N"
	EastOrientation  Orientation = "E"
	SouthOrientation Orientation = "S"
	WestOrientation  Orientation = "W"
)

// Position represents the coordinates (X, Y) of a robot.
type Position struct {
	X, Y int
}

// String returns a string representation of the Position.
func (p *Position) String() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

// CreateRobot creates a new robot with the specified ID and initial position.
func CreateRobot(ID int, input string) (*Robot, error) {
	// Parse the input string to extract initial position and orientation
	parts := strings.Fields(input)
	x, err1 := strconv.Atoi(parts[0])
	y, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("could not parse robot #%d initial position from file", ID)
	}
	// Create and return a new Robot instance
	return &Robot{
		ID:          ID,
		Position:    Position{X: x, Y: y},
		Orientation: Orientation(parts[2]),
	}, nil
}

// Move moves the robot forward in the direction it is facing.
func (r *Robot) Move() {
	switch r.Orientation {
	case NorthOrientation:
		r.Y++
	case EastOrientation:
		r.X++
	case SouthOrientation:
		r.Y--
	case WestOrientation:
		r.X--
	}
}

// CalcPositionAfterMovingForward calculates the position after moving the robot forward.
func (r *Robot) CalcPositionAfterMovingForward() *Position {
	// Create a copy of the original position
	cPosition := &Position{
		X: r.X, Y: r.Y,
	}
	// Update the copied position based on the robot's orientation
	switch r.Orientation {
	case NorthOrientation:
		cPosition.Y += 1
	case EastOrientation:
		cPosition.X += 1
	case SouthOrientation:
		cPosition.Y -= 1
	case WestOrientation:
		cPosition.X -= 1
	}
	return cPosition
}

// TurnLeft rotates the robot 90º to the left.
func (r *Robot) TurnLeft() {
	switch r.Orientation {
	case NorthOrientation:
		r.Orientation = WestOrientation
	case EastOrientation:
		r.Orientation = NorthOrientation
	case SouthOrientation:
		r.Orientation = EastOrientation
	case WestOrientation:
		r.Orientation = SouthOrientation
	}
}

// TurnRight rotates the robot 90º to the right.
func (r *Robot) TurnRight() {
	switch r.Orientation {
	case NorthOrientation:
		r.Orientation = EastOrientation
	case EastOrientation:
		r.Orientation = SouthOrientation
	case SouthOrientation:
		r.Orientation = WestOrientation
	case WestOrientation:
		r.Orientation = NorthOrientation
	}
}
