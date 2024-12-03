package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

func getInputLists() []string {
	reader := bufio.NewReader(os.Stdin)
	output := []string{}
	for true {
		// Just gonna group these into lines for convenience
		// No func call can be split across lines so should be fine
		// to process independently
		text, err := reader.ReadString('\n')
		output = append(output, text)
		if err != nil || err == io.EOF {
			break
		}
	}
	return output
}

type MulState struct {
	NextLetter      rune
	FirstParenSeen  bool
	FirstOperand    string
	CommaSeen       bool
	SecondOperand   string
	SecondParanSeen bool
}

// Returns true if consumed char, false if state is reset
func (state *MulState) ConsumeChar(curChar rune) bool {
	decimals := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	consumed := false

	// If we need to parse "mul"
	if state.NextLetter != 0 {
		// If next char doesn't match, reset mState and continue
		if curChar != state.NextLetter {
			*state = MulState{NextLetter: 'm'}
		} else {
			if curChar == 'm' {
				state.NextLetter = 'u'
			} else if curChar == 'u' {
				state.NextLetter = 'l'
			} else {
				state.NextLetter = 0
			}
			consumed = true
		}
	} else if !state.FirstParenSeen {
		if curChar != '(' {
			*state = MulState{NextLetter: 'm'}
		} else {
			state.FirstParenSeen = true
			consumed = true
		}
	} else if len(state.FirstOperand) < 3 && !state.CommaSeen {
		// We need to keep parsing the first operand
		if curChar == ',' {
			// Need at least one operand char before continuing
			if len(state.FirstOperand) < 1 {
				*state = MulState{NextLetter: 'm'}
			} else {
				state.CommaSeen = true
				consumed = true
			}
		} else if slices.Contains(decimals, curChar) {
			state.FirstOperand += string(curChar)
			consumed = true
		} else {
			// If not decimal or comma then reset state
			*state = MulState{NextLetter: 'm'}
		}
	} else if !state.CommaSeen {
		if curChar != ',' {
			*state = MulState{NextLetter: 'm'}
		} else {
			state.CommaSeen = true
			consumed = true
		}
	} else if len(state.SecondOperand) < 3 {
		if curChar == ')' {
			// Can't end op without second operand
			if len(state.SecondOperand) < 1 {
				*state = MulState{NextLetter: 'm'}
			} else {
				state.SecondParanSeen = true
				consumed = true
			}
		} else if slices.Contains(decimals, curChar) {
			state.SecondOperand += string(curChar)
			consumed = true
		} else {
			*state = MulState{NextLetter: 'm'}
		}
	} else if curChar == ')' {
		// If everything else is set, then we're just looking for the closing paren
		state.SecondParanSeen = true
		consumed = true
	} else {
		// Must be bad state
		*state = MulState{NextLetter: 'm'}
	}
	return consumed
}

type AuthState struct {
	NextLetter rune
	IsDo       bool
}

// Returns true if char is consumed and false if state is reset
func (state *AuthState) ConsumeChar(curChar rune) bool {
	consumed := false
	// Check if this is potentially a "do"
	if state.NextLetter == 'n' && curChar == '(' {
		state.NextLetter = ')'
		state.IsDo = true
		consumed = true
	} else if state.NextLetter == curChar {
		if state.NextLetter == 'd' {
			state.NextLetter = 'o'
		} else if state.NextLetter == 'o' {
			state.NextLetter = 'n'
		} else if state.NextLetter == 'n' {
			state.NextLetter = '\''
		} else if state.NextLetter == '\'' {
			state.NextLetter = 't'
		} else if state.NextLetter == 't' {
			state.NextLetter = '('
		} else if state.NextLetter == '(' {
			state.NextLetter = ')'
		} else if state.NextLetter == ')' {
			state.NextLetter = 0
		}
		consumed = true
	} else {
		*state = AuthState{NextLetter: 'd'}
	}
	return consumed
}

func PartOne() {
	inputLines := getInputLists()
	multState := &MulState{NextLetter: 'm'}
	total := 0
	for _, line := range inputLines {
		for _, char := range line {
			multState.ConsumeChar(char)

			// If mul is complete, evaluate func call
			if multState.SecondParanSeen {
				firstArg, err := strconv.Atoi(multState.FirstOperand)
				if err != nil {
					log.Fatalf("Couldn't parse string from first operand: %s\n", multState.FirstOperand)
				}
				secondArg, err := strconv.Atoi(multState.SecondOperand)
				if err != nil {
					log.Fatalf("Couldn't parse string from first operand: %s\n", multState.SecondOperand)
				}
				total += firstArg * secondArg
				multState = &MulState{NextLetter: 'm'}
			}
		}
	}
	fmt.Printf("Total of mult ops: %d\n", total)
}

func PartTwo() {
	inputLines := getInputLists()
	multState := &MulState{NextLetter: 'm'}
	authState := &AuthState{NextLetter: 'd'}
	shouldMult := true // By default we process mults

	total := 0
	for _, line := range inputLines {
		for _, char := range line {
			// Try to consume with existing mult state
			consumed := multState.ConsumeChar(char)

			// Try to consume with reset mult state
			if !consumed {
				consumed = multState.ConsumeChar(char)
			}

			// Try to consume with existing auth state
			if !consumed {
				consumed = authState.ConsumeChar(char)
			}

			// Try to consume with reset auth state
			if !consumed {
				consumed = authState.ConsumeChar(char)
			}

			// Only one operator can complete at a specific char

			// Check if do/dont state has changed
			if authState.NextLetter == 0 {
				shouldMult = authState.IsDo
				authState = &AuthState{NextLetter: 'd'}
			}

			// If mul is complete, evaluate func call
			if multState.SecondParanSeen {
				if shouldMult {
					firstArg, err := strconv.Atoi(multState.FirstOperand)
					if err != nil {
						log.Fatalf("Couldn't parse string from first operand: %s\n", multState.FirstOperand)
					}
					secondArg, err := strconv.Atoi(multState.SecondOperand)
					if err != nil {
						log.Fatalf("Couldn't parse string from first operand: %s\n", multState.SecondOperand)
					}
					total += firstArg * secondArg
				}
				multState = &MulState{NextLetter: 'm'}
			}
		}
	}
	fmt.Printf("Total of mult ops: %d\n", total)
}

func main() {
	if os.Args[1] == "part_one" {
		PartOne()
	} else if os.Args[1] == "part_two" {
		PartTwo()
	} else {
		fmt.Println("Invalid argument, must be \"part_one\" or \"part_two\"")
	}
}
