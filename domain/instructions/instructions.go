package instructions

import "fmt"

type Instruction string

const (
	RotateRight Instruction = "R"
	RotateLeft  Instruction = "L"
	MoveForward Instruction = "M"
)

type InstructionSet []Instruction

func ParseInstructions(input string) (InstructionSet, error) {
	var res = make([]Instruction, len(input))
	for i, inst := range input {
		switch inst {
		case 'R':
			res[i] = Instruction(RotateRight)
		case 'L':
			res[i] = Instruction(RotateLeft)
		case 'M':
			res[i] = Instruction(MoveForward)
		default:
			return nil, fmt.Errorf("instruction '%c' at position %d is not valid", inst, i)
		}
	}

	return res, nil
}
