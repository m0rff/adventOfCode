package day8

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  int
	ran       bool
}

type Program struct {
	Instructions map[int]Instruction
	accumulator  int
	aborted      bool
}

func Process(data []string) Program {
	p := Program{
		Instructions: map[int]Instruction{},
		accumulator:  0,
		aborted:      false,
	}

	for i, line := range data {
		if strings.TrimSpace(line) == "" {
			p.Instructions[i] = Instruction{
				operation: "END",
				argument:  0,
				ran:       false,
			}
		} else {
			p.Instructions[i] = parseLine(line)
		}
	}

	return p
}

func parseLine(line string) Instruction {
	parts := strings.Split(line, " ")

	return Instruction{
		operation: parts[0],
		argument:  parseArgument(parts[1]),
		ran:       false,
	}
}

func parseArgument(argument string) int {
	value, _ := strconv.ParseInt(argument, 10, 16)

	return int(value)
}

func (p *Program) Run() {
	nextLine := 0

	for {
		instruction := p.Instructions[nextLine]

		if instruction.ran == true {
			p.aborted = true
			break
		}

		if instruction.operation == "END" {
			p.aborted = false
			break
		}

		instruction.ran = true
		p.Instructions[nextLine] = instruction

		switch instruction.operation {
		case "nop":
			nextLine++
		case "acc":
			p.accumulator += instruction.argument
			nextLine++
		case "jmp":
			nextLine += instruction.argument
		}
	}
}

func ChangeOneOperator(program Program, instructionToChange string, lastIndex int) (int, Program) {
	fmt.Println("CHANGING " + instructionToChange + " from index " + strconv.Itoa(lastIndex))
	for i := lastIndex; i < len(program.Instructions); i++ {
		if program.Instructions[i].operation == instructionToChange {
			instruction := program.Instructions[i]
			if instructionToChange == "jmp" {
				instruction.operation = "nop"
			} else {
				instruction.operation = "jmp"
			}
			fmt.Printf("changed instruction from %s to %s at index %d\r\n", instructionToChange, instruction.operation, i)
			program.Instructions[i] = instruction

			return i, program
		}
	}

	return lastIndex, program
}

func CopyProgram(program Program) Program {
	newMap := map[int]Instruction{}
	for k, v := range program.Instructions {
		newMap[k] = v
	}

	return Program{
		accumulator:  0,
		aborted:      false,
		Instructions: newMap,
	}
}

func Find(p Program, instruction string) {
	lastChangedIndex := 0
	for {
		pcopy := CopyProgram(p)
		lastChangedIndex, pcopy = ChangeOneOperator(pcopy, instruction, lastChangedIndex)
		pcopy.aborted = false
		pcopy.Run()
		if pcopy.aborted == false {
			fmt.Println("Program terminated with:")
			fmt.Println(pcopy.accumulator)
			break
		}
		fmt.Println("Program loops. next..")
		lastChangedIndex++
		if lastChangedIndex > len(p.Instructions) {
			break
		}
		pcopy = Program{}
	}
}

func Main(data []string) {
	program := Process(data)
	Find(program, "nop")
	Find(program, "jmp")
}
