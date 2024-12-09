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

func getInputLists() []int {
	reader := bufio.NewReader(os.Stdin)
	output := []int{}
	for true {
		text, err := reader.ReadString('\n')
		if text != "" && text != "\n" {
			for _, char := range strings.TrimRight(text, "\n") {
				val, err := strconv.Atoi(string(char))
				if err != nil {
					log.Fatalf("Failed to parse input char: %c\n", char)
				}
				output = append(output, val)
			}
		}
		if err != nil || err == io.EOF {
			break
		}
	}
	return output
}

func PartOne() {
	input := getInputLists()

	blockIdxs := []int{}
	processingBlock := true
	for i := range input {
		if processingBlock {
			blockIdxs = append(blockIdxs, i)
		}
		processingBlock = !processingBlock
	}

	total := 0
	outputIdx := 0
	headIdx := 0
	tailIdx := len(blockIdxs) - 1
	processingBlock = true
	debugOutput := ""
	for _, count := range input {
		curCount := count
		var curIdx *int
		if processingBlock {
			// When processing files, pull from head Idx
			curIdx = &headIdx
		} else {
			// When processing free space, pull from tail Idx
			curIdx = &tailIdx
		}
		if *curIdx >= len(blockIdxs) {
			// No more blocks
			break
		}

		done := false
		for curCount > 0 {
			for input[blockIdxs[*curIdx]] == 0 {
				if processingBlock {
					*curIdx += 1
					// No more blocks
					if *curIdx >= len(blockIdxs) {
						done = true
						break
					}
				} else {
					*curIdx -= 1
					// No more blocks
					if *curIdx < 0 {
						done = true
						break
					}
				}
			}
			if done {
				break
			}
			total += (*curIdx) * outputIdx
			debugOutput += strconv.Itoa(*curIdx)
			input[blockIdxs[*curIdx]] -= 1
			outputIdx += 1
			curCount -= 1
		}
		if done {
			break
		}
		processingBlock = !processingBlock
	}

	fmt.Printf("Output: %s\n", debugOutput)
	fmt.Printf("Total: %d\n", total)
}

func PartTwo() {

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
