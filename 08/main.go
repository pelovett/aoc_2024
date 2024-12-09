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
		if text != "" && text != "\n" {
			output = append(output, strings.TrimRight(text, "\n"))
		}
		if err != nil || err == io.EOF {
			break
		}
	}
	return output
}

func PartOne() {
	grid := getInputLists()
	for _, row := range grid {
		fmt.Printf("%s\n", row)
	}
	numRow := len(grid)
	numCol := len(grid[0])

	antennas := map[rune][][]int{}
	for r, row := range grid {
		for c, tile := range row {
			if tile != '.' {
				antennas[tile] = append(antennas[tile], []int{r, c})
			}
		}
	}

	antiNodes := map[int]struct{}{}
	for sigType, sources := range antennas {
		fmt.Printf("Processing type: %c\n", sigType)
		// Loop over all pairs to check if in bounds
		for aIdx, a := range sources {
			aRow := a[0]
			aCol := a[1]
			for bIdx, b := range sources {
				if aIdx == bIdx {
					continue
				}
				bRow := b[0]
				bCol := b[1]
				var antiRow int
				var antiCol int
				if aRow > bRow {
					antiRow = aRow + (aRow - bRow)
					// Too low on map
					if antiRow >= numRow {
						continue
					}
				} else if aRow < bRow {
					antiRow = aRow - (bRow - aRow)
					// Too high on map
					if antiRow < 0 {
						continue
					}
				}

				if aCol > bCol {
					antiCol = aCol + (aCol - bCol)
					if antiCol >= numCol {
						continue
					}
				} else if aCol < bCol {
					antiCol = aCol - (bCol - aCol)
					if antiCol < 0 {
						continue
					}
				}
				//fmt.Printf("Found anti node at: %d,%d\n", antiRow, antiCol)
				antiNodes[antiRow*numCol+antiCol] = struct{}{}
			}
		}
	}
	fmt.Printf("Total unique antinodes: %d\n", len(antiNodes))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func PartTwo() {
	grid := getInputLists()
	for _, row := range grid {
		fmt.Printf("%s\n", row)
	}
	numRow := len(grid)
	numCol := len(grid[0])

	antennas := map[rune][][]int{}
	for r, row := range grid {
		for c, tile := range row {
			if tile != '.' {
				antennas[tile] = append(antennas[tile], []int{r, c})
			}
		}
	}

	antiNodes := map[int]struct{}{}
	for sigType, sources := range antennas {
		fmt.Printf("Processing type: %c\n", sigType)
		// Loop over all pairs to check if in bounds
		for aIdx, a := range sources {
			aRow := a[0]
			aCol := a[1]
			for bIdx, b := range sources {
				if aIdx == bIdx {
					continue
				}
				bRow := b[0]
				bCol := b[1]
				var rowDelta int
				var colDelta int
				if aRow > bRow {
					rowDelta = aRow - bRow
				} else if aRow < bRow {
					rowDelta = aRow - bRow
				}

				if aCol > bCol {
					colDelta = aCol - bCol
				} else if aCol < bCol {
					colDelta = aCol - bCol

				}

				fmt.Printf("For pair: %d,%d | %d,%d delta is: %d,%d\n", aRow, aCol, bRow, bCol, rowDelta, colDelta)
				antiNodes[aRow*numCol+aCol] = struct{}{}
				antiNodes[bRow*numCol+bCol] = struct{}{}
				// Try to find smallest delta (one of them must not be 1)
				if abs(rowDelta) != 1 && abs(colDelta) != 1 {
					factor := numCol
					for factor > 1 {
						if rowDelta%factor == 0 && colDelta%factor == 0 {
							break
						}
						factor -= 1
					}

					// If we found a factor, adjust deltas
					if factor != 1 {
						rowDelta /= factor
						colDelta /= factor
						fmt.Printf("\tFound factor: %d new deltas: %d,%d\n", factor, rowDelta, colDelta)
					}
				}

				curRow := aRow + rowDelta
				curCol := aCol + colDelta
				for curRow >= 0 && curRow < numRow && curCol >= 0 && curCol < numCol {
					fmt.Printf("Found antinode at: %d,%d\n", curRow, curCol)
					if grid[curRow][curCol] == '.' {
						grid[curRow] = grid[curRow][:curCol] + string('#') + grid[curRow][curCol+1:]
					}
					antiNodes[curRow*numCol+curCol] = struct{}{}
					curRow += rowDelta
					curCol += colDelta
				}
			}
		}
	}
	for _, row := range grid {
		fmt.Printf("%s\n", row)
	}
	fmt.Printf("Total unique antinodes: %d\n", len(antiNodes))
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
