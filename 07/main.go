package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInputLists() ([]int, [][]int) {
	targetOutput := []int{}
	argsOutput := [][]int{}
	reader := bufio.NewReader(os.Stdin)
	for true {
		text, err := reader.ReadString('\n')
		if text != "\n" && text != "" {
			text = strings.TrimRight(text, "\n")
			splitText := strings.Split(text, ":")
			target, err := strconv.Atoi(splitText[0])
			if err != nil {
				log.Fatalf("Failed to parse target: %s\n", splitText[0])
			}
			targetOutput = append(targetOutput, target)

			args := []int{}
			// Args start with leading space so get slice starting at index 1
			for _, argStr := range strings.Split(splitText[1], " ")[1:] {
				arg, err := strconv.Atoi(argStr)
				if err != nil {
					log.Fatalf("Failed to parse arg: %s\n", argStr)
				}
				args = append(args, arg)
			}
			argsOutput = append(argsOutput, args)
		}
		if err != nil || err == io.EOF {
			break
		}
	}
	return targetOutput, argsOutput
}

func recurseCheckOne(target, cur int, args []int) bool {
	if target == cur && len(args) == 0 {
		return true
	} else if len(args) == 0 {
		return false
	} else if recurseCheckOne(target, cur*args[0], args[1:]) {
		return true
	} else if recurseCheckOne(target, cur+args[0], args[1:]) {
		return true
	} else {
		return false
	}
}

func PartOne() {
	targets, argSlices := getInputLists()
	fmt.Printf("Targets: %v\n", targets)
	fmt.Printf("Args: %v\n", argSlices)

	total := 0
	for i, target := range targets {
		args := argSlices[i]
		if recurseCheckOne(target, 0, args) {
			total += target
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func doubleCheckSolution(target int, solution string) bool {
	args := strings.Split(solution, " ")
	op := ""
	total, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Failed to parse during double check: %s << %s\n", args[0], solution)
	}
	for _, arg := range args[1:] {
		if arg == "+" || arg == "*" || arg == "||" {
			op = arg
			continue
		}

		// Must be next int arg
		cur, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("Failed to parse during double check: %s << %s\n", arg, solution)
		}
		if op == "+" {
			total += cur
		} else if op == "*" {
			total *= cur
		} else if op == "||" {
			total, err = strconv.Atoi(strconv.Itoa(total) + arg)
			if err != nil {
				log.Fatalf("Failed to combine: %d %s << %s\n", total, arg, solution)
			}
		}
	}

	return total == target
}

func recurseCheckTwo(target, cur int, args []int, solution string) bool {
	if target == cur && len(args) == 0 {
		if !doubleCheckSolution(target, solution) {
			log.Fatalf("Solution double check failed!!!!\n")
		}
		return true
	} else if len(args) == 0 {
		return false
	}

	maybeMult := recurseCheckTwo(target, cur*args[0], args[1:], solution+" * "+strconv.Itoa(args[0]))
	if maybeMult {
		return true
	}

	maybeAddition := recurseCheckTwo(target, cur+args[0], args[1:], solution+" + "+strconv.Itoa(args[0]))
	if maybeAddition {
		return true
	}

	combined, err := strconv.Atoi(strconv.Itoa(cur) + strconv.Itoa(args[0]))
	if err != nil {
		log.Fatalf("Failed to parse combined string: %s\n", strconv.Itoa(cur)+strconv.Itoa(args[0]))
	}
	if recurseCheckTwo(target, combined, args[1:], solution+" || "+strconv.Itoa(args[0])) {
		return true
	}
	return false
}

func PartTwo() {
	targets, argSlices := getInputLists()
	fmt.Printf("Targets: %v\n", targets)
	fmt.Printf("Args: %v\n", argSlices)

	total := 0
	numPossible := 0
	for i, target := range targets {
		args := argSlices[i]
		if recurseCheckTwo(target, args[0], args[1:], strconv.Itoa(args[0])) {
			total += target
			numPossible += 1
		}
	}

	fmt.Printf("Possible Solutions: %d\n", numPossible)
	fmt.Printf("Total: %d\n", total)
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
