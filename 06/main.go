package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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
	for _, row := range output {
		fmt.Printf("%s\n", row)
	}
	fmt.Println()
	return output
}

func PartOne() {
	gridMap := getInputLists()

	// Go through grid and mark blocks in columns and rows
	posRow := -1
	posCol := -1
	rowBlocks := make([][]int, len(gridMap))    // In row x, list of col indexes that are blocked
	colBlocks := make([][]int, len(gridMap[0])) // In column x, list of row indexes that are blocked
	numRows := len(gridMap)
	numCols := len(gridMap[0])
	//fmt.Printf("Num row: %d Num Col: %d\n", len(rowBlocks), len(colBlocks))
	for i, row := range gridMap {
		for j, val := range row {
			if val == '#' {
				colBlocks[j] = append(colBlocks[j], i)
				rowBlocks[i] = append(rowBlocks[i], j)
			} else if val == '^' {
				posRow = i
				posCol = j
			}
		}
	}
	fmt.Printf("Start pos: %d, %d\n", posRow, posCol)
	fmt.Printf("Col Blocks:\n%v\n", colBlocks)
	fmt.Printf("\nRow Blocks:\n%v\n", rowBlocks)

	visited := map[int]struct{}{}
	direction := 'U' // Up U | Down D | Left L | Right R
	for true {
		// fmt.Printf("Traveling in direction: %c\n", direction)
		blockedIdx := -1
		// Travel in decreasing row direction
		if direction == 'U' {
			colBlock := colBlocks[posCol]
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, colBlock)
			// Travel in reverse order
			for i := len(colBlock) - 1; i >= 0; i-- {
				// fmt.Printf("Checking value at idx: %d which is %d\n", i, colBlock[i])
				cur := colBlock[i]
				// If a square is blocked in a row earlier then us, we've hit it
				if cur < posRow {
					// fmt.Printf("Blocked!\n")
					blockedIdx = cur + 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				for i := 0; i < posRow; i++ {
					visited[numCols*i+posCol] = struct{}{}
					gridMap[i] = gridMap[i][:posCol] + string('X') + gridMap[i][posCol+1:]
				}
				break
			}

			// Set visited for all tiles between start pos and blocked row
			for i := blockedIdx; i <= posRow; i++ {
				visited[numCols*i+posCol] = struct{}{}
				gridMap[i] = gridMap[i][:posCol] + string('X') + gridMap[i][posCol+1:]
			}
			direction = 'R'
			posRow = blockedIdx
		} else if direction == 'R' {
			rowBlock := rowBlocks[posRow]
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, rowBlock)

			for i := 0; i < len(rowBlock); i++ {
				// fmt.Printf("Checking value at idx: %d which is %d\n", i, rowBlock[i])
				cur := rowBlock[i]
				// If a square is blocked in a col after us, we've hit it
				if posCol < cur {
					// fmt.Printf("Blocked\n")
					blockedIdx = cur - 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				for i := posCol; i < numCols; i++ {
					visited[numCols*posRow+i] = struct{}{}
					gridMap[posRow] = gridMap[posRow][:i] + string('X') + gridMap[posRow][i+1:]
				}
				break
			}

			// Set visited for all tiles between start pos and blocked col
			for i := posCol; i <= blockedIdx; i++ {
				visited[numCols*posRow+i] = struct{}{}
				gridMap[posRow] = gridMap[posRow][:i] + string('X') + gridMap[posRow][i+1:]
			}
			direction = 'D'
			posCol = blockedIdx
		} else if direction == 'D' {
			colBlock := colBlocks[posCol]
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, colBlock)
			// Travel in reverse order
			for i := 0; i < len(colBlock); i++ {
				// fmt.Printf("Checking value at idx: %d which is %d\n", i, colBlock[i])
				cur := colBlock[i]
				// If a square is blocked in a row later then us, we've hit it
				if posRow < cur {
					// fmt.Printf("Blocked!\n")
					blockedIdx = cur - 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				for i := posRow; i < numRows; i++ {
					visited[numCols*i+posCol] = struct{}{}
					gridMap[i] = gridMap[i][:posCol] + string('X') + gridMap[i][posCol+1:]
				}
				break
			}

			// Set visited for all tiles between start pos and blocked row
			for i := posRow; i <= blockedIdx; i++ {
				visited[numCols*i+posCol] = struct{}{}
				gridMap[i] = gridMap[i][:posCol] + string('X') + gridMap[i][posCol+1:]
			}
			direction = 'L'
			posRow = blockedIdx
		} else if direction == 'L' {
			rowBlock := rowBlocks[posRow]
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, rowBlock)

			for i := len(rowBlock) - 1; i >= 0; i-- {
				cur := rowBlock[i]
				// fmt.Printf("Checking value at idx: %d which is %d\n", i, rowBlock[i])

				// If a square is blocked in a col before us, we've hit it
				if cur < posCol {
					// fmt.Printf("Blocked!\n")
					blockedIdx = cur + 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				for i := 0; i <= posCol; i++ {
					visited[numCols*posRow+i] = struct{}{}
					gridMap[posRow] = gridMap[posRow][:i] + string('X') + gridMap[posRow][i+1:]
				}
				break
			}

			// Set visited for all tiles between start pos and blocked col
			for i := blockedIdx; i <= posCol; i++ {
				visited[numCols*posRow+i] = struct{}{}
				gridMap[posRow] = gridMap[posRow][:i] + string('X') + gridMap[posRow][i+1:]
			}
			direction = 'U'
			posCol = blockedIdx
		}
		// for _, row := range gridMap {
		// 	fmt.Printf("%s\n", row)
		// }
		// fmt.Println()
	}

	for _, row := range gridMap {
		fmt.Printf("%s\n", row)
	}
	fmt.Println()
	fmt.Printf("Total visited: %d\n", len(visited))
}

func removeFromSlice(aSlice []int, toRemove int) []int {
	output := []int{}
	for _, val := range aSlice {
		if val != toRemove {
			output = append(output, val)
		}
	}
	return output
}

func checkForCycle(
	direction rune, startRow, startCol int,
	rBlocks, cBlocks [][]int,
	numCols int) bool {

	var visitedUp []int
	var visitedRight []int
	var visitedDown []int
	var visitedLeft []int

	fmt.Printf("Input row blocks: %v\n", rBlocks)
	fmt.Printf("Input col blocks: %v\n", cBlocks)
	rowBlocks := make([][]int, len(rBlocks))
	colBlocks := make([][]int, len(cBlocks))
	for i, row := range rBlocks {
		rowBlocks[i] = make([]int, len(row))
		copy(rowBlocks[i], row)
	}
	for i, col := range cBlocks {
		colBlocks[i] = make([]int, len(col))
		copy(colBlocks[i], col)
	}
	fmt.Printf("Copied row blocks: %v\n", rowBlocks)
	fmt.Printf("Copied col blocks: %v\n", colBlocks)

	// Update block list
	newBlockRow := startRow
	newBlockCol := startCol
	switch direction {
	case 'U':
		newBlockRow -= 1
	case 'R':
		newBlockCol += 1
	case 'D':
		newBlockRow += 1
	case 'L':
		newBlockCol -= 1
	}

	fmt.Printf("\nChecking for cycle...\n")
	fmt.Printf("Currently traveling in dir: %s and at %d,%d and checking block at %d,%d\n", string(direction), startRow, startCol, newBlockRow, newBlockCol)

	// Add new block to our list of blocks
	// fmt.Printf("Before new block:\n")
	// fmt.Printf("Row blocks: %v\n", rowBlocks)
	// fmt.Printf("Col blocks: %v\n", colBlocks)
	if !slices.Contains(rowBlocks[newBlockRow], newBlockCol) {
		rowBlocks[newBlockRow] = append(rowBlocks[newBlockRow], newBlockCol)
		slices.Sort(rowBlocks[newBlockRow])
	}

	if !slices.Contains(colBlocks[newBlockCol], newBlockRow) {
		colBlocks[newBlockCol] = append(colBlocks[newBlockCol], newBlockRow)
		slices.Sort(colBlocks[newBlockCol])
	}
	// fmt.Printf("After new block:\n")
	// fmt.Printf("Row blocks: %v\n", rowBlocks)
	// fmt.Printf("Col blocks: %v\n", colBlocks)

	// Turn from current direction
	var curDir rune
	switch direction {
	case 'U':
		curDir = 'R'
	case 'R':
		curDir = 'D'
	case 'D':
		curDir = 'L'
	case 'L':
		curDir = 'U'
	}

	posRow := startRow
	posCol := startCol
	foundCycle := false
	var blockedIdx int
	for true {
		fmt.Printf("Traveling in direction: %c from pos: %d,%d start: %d,%d\n", curDir, posRow, posCol, startRow, startCol)
		blockedIdx = -1
		// Travel in decreasing row direction
		if curDir == 'U' {
			fmt.Printf("Visited Up: %d\n", len(visitedUp))
			colBlock := colBlocks[posCol]
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, colBlock)
			// Travel in reverse order
			for i := len(colBlock) - 1; i >= 0; i-- {
				// fmt.Printf("Checking value at idx: %d which is %d\n", i, colBlock[i])
				cur := colBlock[i]
				// If a square is blocked in a row earlier then us, we've hit it
				if cur < posRow {
					// fmt.Printf("Blocked!\n")
					blockedIdx = cur + 1
					break
				}
			}

			// If we hit nothing than no cycle exists
			if blockedIdx == -1 {
				fmt.Printf("Leaving maze, no cycle...\n")
				break
			}

			// Set visited for all tiles between start pos and blocked row
			for i := posRow; i > blockedIdx; i-- {
				// If we've ever visited this tile traveling this direction, we're in a cycle
				if slices.Contains(visitedUp, i*numCols+posCol) {
					foundCycle = true
					break
				}
				// if slices.Contains(visitedUp, i*numCols+posCol) {
				// 	foundCycle = true
				// 	break
				// }
				visitedUp = append(visitedUp, i*numCols+posCol)
			}
			if foundCycle {
				break
			}
			curDir = 'R'
			posRow = blockedIdx
		} else if curDir == 'R' {
			fmt.Printf("Visited Right: %d\n", len(visitedRight))
			rowBlock := rowBlocks[posRow]
			//fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, rowBlock)

			for i := 0; i < len(rowBlock); i++ {
				fmt.Printf("Checking value at idx: %d which is %d\n", i, rowBlock[i])
				cur := rowBlock[i]
				// If a square is blocked in a col after us, we've hit it
				if posCol < cur {
					fmt.Printf("Blocked\n")
					blockedIdx = cur - 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				fmt.Printf("Leaving maze, no cycle...\n")
				break
			}

			// Set visited for all tiles between start pos and blocked col
			// fmt.Printf("visitedDown: \n")
			// for _, val := range visitedDown {
			// 	fmt.Printf("%d,%d\n", val/numCols, val%numCols)
			// }
			// fmt.Printf("Staring from %d,%d going to %d,%d\n", posRow, posCol, posRow, blockedIdx)
			for i := posCol; i <= blockedIdx; i++ {
				// If we've ever visited this tile traveling this direction, we're in a cycle
				// if slices.Contains(visitedRight, posRow*numCols+i) {
				// 	foundCycle = true
				// 	break
				// }
				if slices.Contains(visitedRight, posRow*numCols+i) {
					foundCycle = true
					break
				}
				visitedRight = append(visitedRight, posRow*numCols+i)
			}
			if foundCycle {
				break
			}
			curDir = 'D'
			posCol = blockedIdx
		} else if curDir == 'D' {
			fmt.Printf("Visited Down: %d\n", len(visitedDown))
			colBlock := colBlocks[posCol]
			//fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, colBlock)
			// Travel in reverse order
			for i := 0; i < len(colBlock); i++ {
				//fmt.Printf("Checking value at idx: %d which is %d\n", i, colBlock[i])
				cur := colBlock[i]
				// If a square is blocked in a row later then us, we've hit it
				if posRow < cur {
					blockedIdx = cur - 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				fmt.Printf("Leaving maze, no cycle...\n")
				break
			}

			// Set visited for all tiles between start pos and blocked row
			for i := posRow; i <= blockedIdx; i++ {
				// At each position, if we took a right turn, what whould we hit a block
				// if slices.Contains(visitedDown, i*numCols+posCol) {
				// 	foundCycle = true
				// 	break
				// }
				if slices.Contains(visitedDown, i*numCols+posCol) {
					foundCycle = true
					break
				}
				visitedDown = append(visitedDown, i*numCols+posCol)
			}
			if foundCycle {
				break
			}
			curDir = 'L'
			posRow = blockedIdx
		} else if curDir == 'L' {
			fmt.Printf("Visited Left: %d\n", len(visitedLeft))
			rowBlock := rowBlocks[posRow]
			// fmt.Printf("Row blocks: %v\n", rowBlocks)
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, rowBlock)

			for i := len(rowBlock) - 1; i >= 0; i-- {
				cur := rowBlock[i]
				//fmt.Printf("Checking value at idx: %d which is %d\n", i, rowBlock[i])

				// If a square is blocked in a col before us, we've hit it
				if cur < posCol {
					fmt.Printf("Blocked!\n")
					blockedIdx = cur + 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				fmt.Printf("Leaving maze, no cycle...\n")
				break
			}

			// Set visited for all tiles between start pos and blocked col
			for i := blockedIdx; i <= posCol; i++ {
				// At each position, if we took a right turn, what whould we hit a block
				// if slices.Contains(visitedLeft, posRow*numCols+i) {
				// 	foundCycle = true
				// 	break
				// }
				if slices.Contains(visitedLeft, posRow*numCols+i) {
					foundCycle = true
					break
				}
				visitedLeft = append(visitedLeft, posRow*numCols+i)
			}
			if foundCycle {
				break
			}
			curDir = 'U'
			posCol = blockedIdx
		}
	}
	return foundCycle
}

func PartTwo() {
	gridMap := getInputLists()

	// Go through grid and mark blocks in columns and rows
	posRow := -1
	posCol := -1
	rowBlocks := make([][]int, len(gridMap))    // In row x, list of col indexes that are blocked
	colBlocks := make([][]int, len(gridMap[0])) // In column x, list of row indexes that are blocked
	numRows := len(gridMap)
	numCols := len(gridMap[0])
	//fmt.Printf("Num row: %d Num Col: %d\n", len(rowBlocks), len(colBlocks))
	for i, row := range gridMap {
		for j, val := range row {
			if val == '#' {
				colBlocks[j] = append(colBlocks[j], i)
				rowBlocks[i] = append(rowBlocks[i], j)
			} else if val == '^' {
				posRow = i
				posCol = j
			}
		}
	}
	startRow := posRow
	startCol := posCol
	fmt.Printf("Start pos: %d, %d\n", posRow, posCol)
	fmt.Printf("Col Blocks:\n%v\n", colBlocks)
	fmt.Printf("\nRow Blocks:\n%v\n", rowBlocks)

	visitedUp := []int{}
	visitedRight := []int{}
	visitedDown := []int{}
	visitedLeft := []int{}

	total := 0
	potentialBlocks := map[int]struct{}{}
	direction := 'U' // Up U | Down D | Left L | Right R
	for true {
		fmt.Printf("Traveling in direction: %c\n", direction)
		blockedIdx := -1
		// Travel in decreasing row direction
		if direction == 'U' {
			colBlock := colBlocks[posCol]
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, colBlock)
			// Travel in reverse order
			for i := len(colBlock) - 1; i >= 0; i-- {
				// fmt.Printf("Checking value at idx: %d which is %d\n", i, colBlock[i])
				cur := colBlock[i]
				// If a square is blocked in a row earlier then us, we've hit it
				if cur < posRow {
					// fmt.Printf("Blocked!\n")
					blockedIdx = cur + 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				for i := posRow; i > 0; i-- {
					// At each position, if we took a right turn, what whould we hit
					if slices.Contains(visitedRight, i*numCols+posCol) ||
						checkForCycle(direction, i, posCol, rowBlocks, colBlocks, numCols) {
						fmt.Printf("Found potential block if we turn right at row: %d\n", i)
						total += 1
						potentialBlocks[(i-1)*numCols+posCol] = struct{}{}
					}
					gridMap[i] = gridMap[i][:posCol] + string('X') + gridMap[i][posCol+1:]
				}
				break
			}

			// Set visited for all tiles between start pos and blocked row
			for i := posRow; i > blockedIdx; i-- {
				// At each position, if we took a right turn, what whould we hit
				if slices.Contains(visitedRight, i*numCols+posCol) ||
					checkForCycle(direction, i, posCol, rowBlocks, colBlocks, numCols) {
					fmt.Printf("Found potential block if we turn right at row: %d\n", i)
					total += 1
					potentialBlocks[(i-1)*numCols+posCol] = struct{}{}
				}
				visitedUp = append(visitedUp, i*numCols+posCol)
				gridMap[i] = gridMap[i][:posCol] + string('X') + gridMap[i][posCol+1:]
			}
			direction = 'R'
			posRow = blockedIdx
		} else if direction == 'R' {
			rowBlock := rowBlocks[posRow]
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, rowBlock)

			for i := 0; i < len(rowBlock); i++ {
				// fmt.Printf("Checking value at idx: %d which is %d\n", i, rowBlock[i])
				cur := rowBlock[i]
				// If a square is blocked in a col after us, we've hit it
				if posCol < cur {
					// fmt.Printf("Blocked\n")
					blockedIdx = cur - 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				for i := posCol; i < numCols; i++ {
					// At each position, if we took a right turn, what whould we hit a block
					if slices.Contains(visitedDown, posRow*numCols+i) ||
						checkForCycle(direction, posRow, i, rowBlocks, colBlocks, numCols) {
						fmt.Printf("Found potential block if we turn down at col: %d\n", i)
						total += 1
						potentialBlocks[posRow*numCols+i+1] = struct{}{}
					}
					gridMap[posRow] = gridMap[posRow][:i] + string('X') + gridMap[posRow][i+1:]
				}
				break
			}

			// Set visited for all tiles between start pos and blocked col
			for i := posCol; i < blockedIdx; i++ {
				// At each position, if we took a right turn, what whould we hit a block
				if slices.Contains(visitedDown, posRow*numCols+i) ||
					checkForCycle(direction, posRow, i, rowBlocks, colBlocks, numCols) {
					fmt.Printf("Found potential block if we turn down at col: %d\n", i)
					total += 1
					potentialBlocks[posRow*numCols+i+1] = struct{}{}
				}
				visitedRight = append(visitedRight, posRow*numCols+i)
				gridMap[posRow] = gridMap[posRow][:i] + string('X') + gridMap[posRow][i+1:]
			}
			direction = 'D'
			posCol = blockedIdx
		} else if direction == 'D' {
			colBlock := colBlocks[posCol]
			fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, colBlock)
			// Travel in reverse order
			for i := 0; i < len(colBlock); i++ {
				fmt.Printf("Checking value at idx: %d which is %d\n", i, colBlock[i])
				cur := colBlock[i]
				// If a square is blocked in a row later then us, we've hit it
				if posRow < cur {
					fmt.Printf("Blocked!\n")
					blockedIdx = cur - 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				for i := posRow; i < numRows-1; i++ {
					// At each position, if we took a right turn, what whould we hit a block
					if slices.Contains(visitedLeft, i*numCols+posCol) ||
						checkForCycle(direction, i, posCol, rowBlocks, colBlocks, numCols) {
						fmt.Printf("Found potential block if we turn left at row: %d\n", i)
						total += 1
						potentialBlocks[(i+1)*numCols+posCol] = struct{}{}
					}
					gridMap[i] = gridMap[i][:posCol] + string('X') + gridMap[i][posCol+1:]
				}
				break
			}

			// Set visited for all tiles between start pos and blocked row
			for i := posRow; i < blockedIdx; i++ {
				// At each position, if we took a right turn, what whould we hit a block
				if slices.Contains(visitedLeft, i*numCols+posCol) ||
					checkForCycle(direction, i, posCol, rowBlocks, colBlocks, numCols) {
					fmt.Printf("Found potential block if we turn left at row: %d\n", i)
					total += 1
					potentialBlocks[(i+1)*numCols+posCol] = struct{}{}
				}
				visitedDown = append(visitedDown, i*numCols+posCol)
				gridMap[i] = gridMap[i][:posCol] + string('X') + gridMap[i][posCol+1:]
			}
			direction = 'L'
			posRow = blockedIdx
		} else if direction == 'L' {
			rowBlock := rowBlocks[posRow]
			// fmt.Printf("Current pos: %d,%d looking for blocks in list: %v\n", posRow, posCol, rowBlock)

			for i := len(rowBlock) - 1; i >= 0; i-- {
				cur := rowBlock[i]
				// fmt.Printf("Checking value at idx: %d which is %d\n", i, rowBlock[i])

				// If a square is blocked in a col before us, we've hit it
				if cur < posCol {
					// fmt.Printf("Blocked!\n")
					blockedIdx = cur + 1
					break
				}
			}

			// If we hit nothing than the guard has left the maze
			if blockedIdx == -1 {
				for i := posCol; i > 0; i-- {
					// At each position, if we took a right turn, what whould we hit a block
					if slices.Contains(visitedUp, posRow*numCols+i) ||
						checkForCycle(direction, posRow, i, rowBlocks, colBlocks, numCols) {
						fmt.Printf("Found potential block if we turn up (for what) at col: %d\n", i)
						total += 1
						potentialBlocks[posRow*numCols+i-1] = struct{}{}
					}
					gridMap[posRow] = gridMap[posRow][:i] + string('X') + gridMap[posRow][i+1:]
				}
				break
			}

			// Set visited for all tiles between start pos and blocked col
			for i := posCol; i >= blockedIdx; i-- {
				// At each position, if we took a right turn, what whould we hit a block
				if slices.Contains(visitedUp, posRow*numCols+i) ||
					checkForCycle(direction, posRow, i, rowBlocks, colBlocks, numCols) {
					fmt.Printf("Found potential block if we turn up (for what) at col: %d\n", i)
					total += 1
					potentialBlocks[posRow*numCols+i-1] = struct{}{}
				}
				visitedLeft = append(visitedLeft, posRow*numCols+i)
				gridMap[posRow] = gridMap[posRow][:i] + string('X') + gridMap[posRow][i+1:]
			}
			direction = 'U'
			posCol = blockedIdx
		}

		for block, _ := range potentialBlocks {
			row := block / numCols
			col := block % numCols
			gridMap[row] = gridMap[row][:col] + string('O') + gridMap[row][col+1:]
		}
		gridMap[posRow] = gridMap[posRow][:posCol] + string('@') + gridMap[posRow][posCol+1:]
		for _, row := range gridMap {
			fmt.Printf("%s\n", row)
		}
		fmt.Println()
	}

	for block, _ := range potentialBlocks {
		row := block / numCols
		col := block % numCols
		gridMap[row] = gridMap[row][:col] + string('O') + gridMap[row][col+1:]
	}

	for _, row := range gridMap {
		fmt.Printf("%s\n", row)
	}
	fmt.Println()

	delete(potentialBlocks, startRow*numCols+startCol)
	fmt.Printf("Potential blocks: %v\n", potentialBlocks)
	fmt.Printf("Num potential blocks: %d\n", len(potentialBlocks))
	fmt.Printf("Total cycles possible: %d\n", total)
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
