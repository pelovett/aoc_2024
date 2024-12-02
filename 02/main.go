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
	output := [][]int{}
	reader := bufio.NewReader(os.Stdin)
	for true {
		inputText, err := reader.ReadString('\n')
		if inputText == "" {
			break
		}
		splitText := strings.Split(strings.TrimRight(inputText, "\n"), " ")
		temp := []int{}
		for _, token := range splitText {
			num, err := strconv.Atoi(token)
			if err != nil {
				fmt.Printf("Error parsing line: '%s'\n", inputText)
				fmt.Printf("Error parsing: '%s'\n", token)
				log.Fatal(err)
			}
			temp = append(temp, num)
		}
		output = append(output, temp)
		if err != nil || err == io.EOF {
			break
		}
	}
	return output
}

func PartOne() {
	input := getInputLists()
	total := 0
	for _, reactorLevels := range input {
		prevLevel := -1
		isIncreasing := false
		increasingFlagSet := false
		badLevels := false
		for _, curLevel := range reactorLevels {
			if prevLevel == -1 {
				prevLevel = curLevel
				continue
			}

			if !increasingFlagSet {
				if prevLevel == curLevel {
					badLevels = true
					break
				} else if prevLevel > curLevel {
					if prevLevel > curLevel+3 {
						badLevels = true
						break
					}
					isIncreasing = false
					increasingFlagSet = true
				} else {
					if prevLevel < curLevel-3 {
						badLevels = true
						break
					}
					isIncreasing = true
					increasingFlagSet = true
				}
			} else {
				if prevLevel == curLevel {
					badLevels = true
					break
				} else if !isIncreasing && (prevLevel < curLevel || prevLevel > curLevel+3) {
					badLevels = true
					break
				} else if isIncreasing && (prevLevel > curLevel || prevLevel < curLevel-3) {
					badLevels = true
					break
				}
			}
			prevLevel = curLevel
		}
		if !badLevels {
			total += 1
		}
	}

	fmt.Printf("Number of good levels: %d\n", total)
}

func PartTwo() {
	input := getInputLists()
	total := 0
	for _, reactorLevels := range input {
		prevLevel := -1
		isIncreasing := false
		increasingFlagSet := false
		badIndex := -1
		for idx, curLevel := range reactorLevels {
			if prevLevel == -1 {
				prevLevel = curLevel
				continue
			}

			if !increasingFlagSet {
				if prevLevel == curLevel {
					badIndex = idx
					break
				} else if prevLevel > curLevel {
					if prevLevel > curLevel+3 {
						badIndex = idx
						break
					}
					isIncreasing = false
					increasingFlagSet = true
				} else {
					if prevLevel < curLevel-3 {
						badIndex = idx
						break
					}
					isIncreasing = true
					increasingFlagSet = true
				}
			} else {
				if prevLevel == curLevel {
					badIndex = idx
					break
				} else if !isIncreasing && (prevLevel < curLevel || prevLevel > curLevel+3) {
					badIndex = idx
					break
				} else if isIncreasing && (prevLevel > curLevel || prevLevel < curLevel-3) {
					badIndex = idx
					break
				}
			}
			prevLevel = curLevel
		}
		if badIndex == -1 {
			total += 1
			continue
		}

		// Try checking levels but with bad index and two previous removed one at a time
		// Use min to not underflow at start of array
		for i := 0; i < min(3, badIndex+1); i++ {
			ignoreIdx := badIndex - i

			prevLevel = -1
			isIncreasing = false
			increasingFlagSet = false
			newBad := -1
			for idx, curLevel := range reactorLevels {
				if idx == ignoreIdx {
					continue
				}

				if prevLevel == -1 {
					prevLevel = curLevel
					continue
				}

				if !increasingFlagSet {
					if prevLevel == curLevel {
						newBad = idx
						break
					} else if prevLevel > curLevel {
						if prevLevel > curLevel+3 {
							newBad = idx
							break
						}
						isIncreasing = false
						increasingFlagSet = true
					} else {
						if prevLevel < curLevel-3 {
							newBad = idx
							break
						}
						isIncreasing = true
						increasingFlagSet = true
					}
				} else {
					if prevLevel == curLevel {
						newBad = idx
						break
					} else if !isIncreasing && (prevLevel < curLevel || prevLevel > curLevel+3) {
						newBad = idx
						break
					} else if isIncreasing && (prevLevel > curLevel || prevLevel < curLevel-3) {
						newBad = idx
						break
					}
				}
				prevLevel = curLevel
			}

			if newBad == -1 {
				total += 1
				break
			}
		}
	}

	fmt.Printf("Number of good levels: %d\n", total)
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
