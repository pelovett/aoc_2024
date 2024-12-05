package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getInputLists() []string {
	reader := bufio.NewReader(os.Stdin)
	output := []string{}
	for true {
		text, err := reader.ReadString('\n')
		if strings.TrimRight(text, "\n") != "" {
			output = append(output, strings.TrimRight(text, "\n"))
		}
		if err != nil || err == io.EOF {
			break
		}
	}
	return output
}

type Node struct {
	row int
	col int
}

func CreateValidDirectionCheckP1(r, c, maxR, maxC int) func(*Node) bool {
	return func(cur *Node) bool {
		//fmt.Printf("Checking %d,%d with %d,%d\n", cur.row, cur.col, r, c)
		if r == -1 && cur.row == 0 {
			return false
		} else if c == -1 && cur.col == 0 {
			return false
		} else if r == 1 && cur.row == maxR {
			return false
		} else if c == 1 && cur.col == maxC {
			return false
		} else {
			return true
		}
	}
}

func PartOne() {
	inputs := getInputLists()
	numRows := len(inputs)
	numCols := len(inputs[0])

	for _, line := range inputs {
		fmt.Println(line)
	}
	fmt.Println()

	reverseIndex := make([]*Node, numRows*numCols)
	xNodes := []*Node{}
	for i, line := range inputs {
		for j, char := range line {
			// Either we've visited before or its empty
			cur := reverseIndex[i*numRows+j]
			if cur == nil {
				cur = &Node{row: i, col: j}
				reverseIndex[i*numRows+j] = cur
			}

			if char == 'X' {
				xNodes = append(xNodes, cur)
			}
		}
	}

	total := 0
	compass := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	for _, xNode := range xNodes {
		for _, direction := range compass {
			rowDiff := direction[0]
			colDiff := direction[1]
			checkValidDirection := CreateValidDirectionCheckP1(rowDiff, colDiff, len(inputs[0])-1, len(inputs)-1)

			if !checkValidDirection(xNode) {
				continue
			}

			if inputs[xNode.row+rowDiff][xNode.col+colDiff] != 'M' {
				continue
			}

			mIdx := (xNode.row+rowDiff)*numRows + xNode.col + colDiff
			mNode := reverseIndex[mIdx]
			if !checkValidDirection(mNode) {
				continue
			}

			if inputs[mNode.row+rowDiff][mNode.col+colDiff] != 'A' {
				continue
			}

			aIdx := (mNode.row+rowDiff)*numRows + mNode.col + colDiff
			aNode := reverseIndex[aIdx]
			if !checkValidDirection(aNode) {
				continue
			}

			if inputs[aNode.row+rowDiff][aNode.col+colDiff] == 'S' {
				total += 1
			}

		}
	}

	fmt.Printf("Total XMASs found: %d\n", total)
}

func PartTwo() {
	inputs := getInputLists()
	numRows := len(inputs)
	numCols := len(inputs[0])

	for _, line := range inputs {
		fmt.Println(line)
	}
	fmt.Println()

	aNodes := []*Node{}
	for i, line := range inputs {
		for j, char := range line {
			if char == 'A' {
				aNodes = append(aNodes, &Node{row: i, col: j})
			}
		}
	}

	total := 0
	deltas := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	for _, aNode := range aNodes {
		if aNode.row == 0 || aNode.row == numRows-1 {
			continue
		} else if aNode.col == 0 || aNode.col == numCols-1 {
			continue
		}

		// Start delta rotation at each possible starting pos
		for i := range deltas {
			// First M
			delta := deltas[i%len(deltas)]
			if inputs[aNode.row+delta[0]][aNode.col+delta[1]] != 'M' {
				continue
			}

			// Second M
			delta = deltas[(i+1)%len(deltas)]
			if inputs[aNode.row+delta[0]][aNode.col+delta[1]] != 'M' {
				continue
			}

			// First S
			delta = deltas[(i+2)%len(deltas)]
			if inputs[aNode.row+delta[0]][aNode.col+delta[1]] != 'S' {
				continue
			}

			// Second S
			delta = deltas[(i+3)%len(deltas)]
			if inputs[aNode.row+delta[0]][aNode.col+delta[1]] != 'S' {
				continue
			}
			total += 1
		}
	}

	fmt.Printf("Total XMASs found: %d\n", total)
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
