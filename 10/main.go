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

func getInputLists() [][]int {
	reader := bufio.NewReader(os.Stdin)
	output := [][]int{}
	for true {
		text, err := reader.ReadString('\n')
		if text != "" && text != "\n" {
			temp := []int{}
			for _, char := range strings.TrimRight(text, "\n") {
				num, err := strconv.Atoi(string(char))
				if err != nil {
					log.Fatalf("Failed to parse while reading input: %c\n", char)
				}
				temp = append(temp, num)
			}
			output = append(output, temp)
		}
		if err != nil || err == io.EOF {
			break
		}
	}
	return output
}

func removeDup(a, b []int) []int {
	output := []int{}
	seen := map[int]struct{}{}
	for _, val := range a {
		if _, exists := seen[val]; !exists {
			output = append(output, val)
			seen[val] = struct{}{}
		}
	}

	for _, val := range b {
		if _, exists := seen[val]; !exists {
			output = append(output, val)
			seen[val] = struct{}{}
		}
	}
	return output
}

func partOneSearch(row, col int, gridPtr *[][]int, knownTilesPtr *map[int][]int, visitedPtr *map[int]struct{}) []int {
	grid := *gridPtr
	knownTiles := *knownTilesPtr
	visited := *visitedPtr
	numCol := len(grid[0])
	intIdx := row*numCol + col

	// If we already visited, return early
	if _, exists := visited[intIdx]; exists {
		return knownTiles[intIdx]
	}

	// If we're at a peak, return
	if grid[row][col] == 9 {
		knownTiles[intIdx] = []int{intIdx}
		visited[intIdx] = struct{}{}
		return []int{intIdx}
	}

	// Otherwise climb
	reachable := []int{}
	nextVal := grid[row][col] + 1
	if row > 0 && grid[row-1][col] == nextVal {
		reachable = removeDup(
			reachable,
			partOneSearch(row-1, col, gridPtr, knownTilesPtr, visitedPtr),
		)
	}

	if row < len(grid)-1 && grid[row+1][col] == nextVal {
		reachable = removeDup(
			reachable,
			partOneSearch(row+1, col, gridPtr, knownTilesPtr, visitedPtr),
		)
	}

	if col > 0 && grid[row][col-1] == nextVal {
		reachable = removeDup(
			reachable,
			partOneSearch(row, col-1, gridPtr, knownTilesPtr, visitedPtr),
		)
	}

	if col < numCol-1 && grid[row][col+1] == nextVal {
		reachable = removeDup(
			reachable,
			partOneSearch(row, col+1, gridPtr, knownTilesPtr, visitedPtr),
		)
	}

	knownTiles[intIdx] = reachable
	visited[intIdx] = struct{}{}
	return reachable
}

func PartOne() {
	grid := getInputLists()

	// Scan grid for all zeros
	zeros := [][]int{}
	for i, row := range grid {
		for j, tile := range row {
			fmt.Printf("%d ", tile)
			if tile == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nFound zeros: %v\n", zeros)

	pointsAfterTile := map[int][]int{}
	visitedTiles := map[int]struct{}{}
	total := 0
	for _, start := range zeros {
		total += len(partOneSearch(start[0], start[1], &grid, &pointsAfterTile, &visitedTiles))
	}
	fmt.Printf("Total: %d\n", total)
}

func concat(a, b []int) []int {
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}

	output := make([]int, len(a)+len(b))
	for i, val := range a {
		output[i] = val
	}
	for i, val := range b {
		output[i+len(a)] = val
	}
	return output
}

func partTwoSearch(row, col int, gridPtr *[][]int, knownTilesPtr *map[int][]int, visitedPtr *map[int]struct{}) []int {
	grid := *gridPtr
	knownTiles := *knownTilesPtr
	visited := *visitedPtr
	numCol := len(grid[0])
	intIdx := row*numCol + col

	// If we already visited, return early
	if _, exists := visited[intIdx]; exists {
		results, _ := knownTiles[intIdx]
		return results
	}

	// If we're at a peak, return
	if grid[row][col] == 9 {
		knownTiles[intIdx] = []int{intIdx}
		visited[intIdx] = struct{}{}
		return []int{intIdx}
	}

	// Otherwise climb
	reachable := []int{}
	nextVal := grid[row][col] + 1
	if row > 0 && grid[row-1][col] == nextVal {
		reachable = concat(
			reachable,
			partTwoSearch(row-1, col, gridPtr, knownTilesPtr, visitedPtr),
		)
	}

	if row < len(grid)-1 && grid[row+1][col] == nextVal {
		reachable = concat(
			reachable,
			partTwoSearch(row+1, col, gridPtr, knownTilesPtr, visitedPtr),
		)
	}

	if col > 0 && grid[row][col-1] == nextVal {
		reachable = concat(
			reachable,
			partTwoSearch(row, col-1, gridPtr, knownTilesPtr, visitedPtr),
		)
	}

	if col < numCol-1 && grid[row][col+1] == nextVal {
		reachable = concat(
			reachable,
			partTwoSearch(row, col+1, gridPtr, knownTilesPtr, visitedPtr),
		)
	}

	knownTiles[intIdx] = reachable
	visited[intIdx] = struct{}{}
	return reachable
}

func PartTwo() {
	grid := getInputLists()

	// Scan grid for all zeros
	zeros := [][]int{}
	for i, row := range grid {
		for j, tile := range row {
			fmt.Printf("%d ", tile)
			if tile == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nFound zeros: %v\n", zeros)

	pointsAfterTile := map[int][]int{}
	visitedTiles := map[int]struct{}{}
	total := 0
	for _, start := range zeros {
		total += len(partTwoSearch(start[0], start[1], &grid, &pointsAfterTile, &visitedTiles))
	}
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
