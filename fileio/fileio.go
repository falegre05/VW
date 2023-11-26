package fileio

import (
	"bufio"
	"os"
	"vw/domain/grid"
	"vw/domain/instructions"
	"vw/domain/robot"
)

func ReadInstructionsFromFile(filePath string) (*grid.Grid, []robot.Robot, []instructions.InstructionSet, error) {
	var (
		g        *grid.Grid
		robots   []robot.Robot
		instSets []instructions.InstructionSet
		err      error
	)

	// Try to open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	robotID := 0
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		switch {
		case i == 0: // Create grid
			g, err = grid.CreateGrid(line)
			if err != nil {
				return nil, nil, nil, err
			}
		case i%2 == 1: // Create robots
			r, err := robot.CreateRobot(robotID, line)
			if err != nil {
				return nil, nil, nil, err
			}
			robots = append(robots, *r)
			robotID++
		case i%2 == 0: // Create instructionSet
			instSet, err := instructions.ParseInstructions(line)
			if err != nil {
				return nil, nil, nil, err
			}
			instSets = append(instSets, instSet)

		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, nil, nil, err
	}

	return g, robots, instSets, nil
}
