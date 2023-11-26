package main

import (
	"fmt"
	"os"
	"vw/domain/service"
	"vw/fileio"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./program.go instructions.txt")
		return
	}
	// Extract the filename from command-line arguments
	filePath := os.Args[1]

	g, robots, instructionSets, err := fileio.ReadInstructionsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	var rService = service.NewRobotService(g)

	for i := range robots {
		if err := rService.ApplyInstructions(&robots[i], instructionSets[i]); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
