package grid

import (
	"errors"
	"strconv"
	"strings"
	"vw/domain/robot"
)

// Grid represents the domain model for the floor grid where robots move.
type Grid struct {
	XMax, YMax int                     // Maximum X and Y coordinates of the grid
	Robots     map[string]*robot.Robot // Map to track robots on the grid using their positions as keys
}

// CreateGrid creates a new grid with the specified maximum X and Y coordinates.
func CreateGrid(input string) (*Grid, error) {
	// Parse the input string to extract maximum grid size
	parts := strings.Fields(input)
	x, err1 := strconv.Atoi(parts[0])
	y, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return nil, errors.New("could not parse grid max size from file")
	}

	// Create and return a new Grid instance with an initialized robot map
	return &Grid{XMax: x, YMax: y, Robots: make(map[string]*robot.Robot)}, nil
}
