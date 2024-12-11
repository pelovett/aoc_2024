package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	_ "net/http/pprof"
)

func getInputLists() []int {
	reader := bufio.NewReader(os.Stdin)
	output := []int{}
	for true {
		text, err := reader.ReadString('\n')
		if text != "" && text != "\n" {
			temp := strings.Split(strings.TrimRight(text, "\n"), " ")
			for _, char := range temp {
				val, err := strconv.Atoi(string(char))
				if err != nil {
					log.Fatalf("failed to parse: %c\n", char)
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

func hasEvenNumDigits(a int) bool {
	return int64(math.Floor(math.Log10(float64(a))))%2 == 1
}

func PartOne() {
	stones := getInputLists()
	fmt.Printf("Input: %v\n", stones)

	step := 1
	for step < 26 {
		temp := []int{}
		for i, stone := range stones {
			if stone == 0 {
				//fmt.Printf("\t0 stone: %d\n", stone)
				stones[i] = 1
			} else if hasEvenNumDigits(stone) {
				//fmt.Printf("\tEven num digits: %d\n", stone)
				strStone := strconv.Itoa(stone)
				firstHalf, err := strconv.Atoi(strStone[:len(strStone)/2])
				if err != nil {
					log.Fatalf("Failed to parse half of stone: %s  half: %s\n", strStone, strStone[:len(strStone)/2])
				}
				secondHalf, err := strconv.Atoi(strStone[len(strStone)/2:])
				if err != nil {
					log.Fatalf("Failed to parse half of stone: %s  half: %s\n", strStone, strStone[len(strStone)/2:])
				}
				stones[i] = firstHalf
				temp = append(temp, secondHalf)
			} else {
				stones[i] *= 2024
				//fmt.Printf("\tDefault mult 2024: %d\n", stone)
			}
		}
		stones = slices.Concat(stones, temp)
		fmt.Printf("Total number of stones: %d\n", len(stones))

		//fmt.Printf("Stones: %v\n", stones)
		step += 1
	}

	// fmt.Printf("Total number of stones: %d\n", len(stones))
}

func PartTwo() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	inputStones := getInputLists()
	fmt.Printf("Input: %v\n", inputStones)

	// Create large arrays to hold stones
	temp := make([]int, 1_000_000)
	stones := []*[]int{&temp}

	rowIdx := 0
	colIdx := 0
	for _, stone := range inputStones {
		(*stones[rowIdx])[colIdx] = stone
		colIdx += 1
	}

	fmt.Printf("Starting rowidx colidx: %d %d\n", rowIdx, colIdx)
	step := 1
	for step < 76 {
		fmt.Printf("%s Starting step: %d\n", time.Now().UTC(), step)
		startRowIdx := rowIdx
		startColIdx := colIdx
		for i, rowPtr := range stones {
			for j, stone := range *rowPtr {
				// fmt.Printf("i j: %d %d | rowidx colidx: %d %d\n", i, j, rowIdx, colIdx)

				// Check if we're at the end of filled array
				if i == startRowIdx && j == startColIdx {
					break
				}
				if stone == 0 {
					// fmt.Printf("\t0 stone: %d\n", stone)
					(*stones[i])[j] = 1
				} else if hasEvenNumDigits(stone) {
					// fmt.Printf("\tEven num digits: %d\n", stone)
					strStone := strconv.Itoa(stone)
					firstHalf, err := strconv.Atoi(strStone[:len(strStone)/2])
					if err != nil {
						log.Fatalf("Failed to parse half of stone: %s  half: %s\n", strStone, strStone[:len(strStone)/2])
					}
					secondHalf, err := strconv.Atoi(strStone[len(strStone)/2:])
					if err != nil {
						log.Fatalf("Failed to parse half of stone: %s  half: %s\n", strStone, strStone[len(strStone)/2:])
					}
					(*stones[i])[j] = firstHalf
					(*stones[rowIdx])[colIdx] = secondHalf
					colIdx += 1
					// Check if we need another big array
					if colIdx == len(*stones[rowIdx]) {
						temp := make([]int, 1_000_000)
						stones = append(stones, &temp)
						rowIdx += 1
						colIdx = 0
					}
				} else {
					(*stones[i])[j] *= 2024
					// fmt.Printf("\tDefault mult 2024: %d\n", stone)
				}
			}
			// fmt.Printf("Row index: %d Col Idx: %d\n", rowIdx, colIdx)
			// fmt.Printf("Total number of stones: %d\n", rowIdx*len(*stones[0])+(colIdx))

			// for i, arrPtr := range stones {
			// 	if rowIdx == i {
			// 		fmt.Printf("Stones: %v\n", (*arrPtr)[:colIdx])
			// 	} else {
			// 		fmt.Printf("Stones: %v\n", *arrPtr)
			// 	}
			// }
		}
		step += 1
	}

	fmt.Printf("Row index: %d Col Idx: %d\n", rowIdx, colIdx)
	fmt.Printf("Total number of stones: %d\n", rowIdx*len(*stones[0])+(colIdx))
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
