package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInputLists() ([]int, []int) {
	reader := bufio.NewReader(os.Stdin)
	leftArr := []int{}
	rightArr := []int{}
	for true {
		text, err := reader.ReadString('\n')
		if text != "" {
			textVal := strings.Split(strings.TrimRight(text, "\n"), "   ")
			leftVal, _ := strconv.Atoi(textVal[0])
			rightVal, _ := strconv.Atoi(textVal[1])
			leftArr = append(leftArr, leftVal)
			rightArr = append(rightArr, rightVal)
		}
		if err != nil || err == io.EOF {
			break
		}
	}
	return leftArr, rightArr
}

func PartOne() {
	leftArr, rightArr := getInputLists()
	totalDiff := 0
	sort.Ints(leftArr)
	sort.Ints(rightArr)
	for i := 0; i < len(leftArr); i++ {
		if leftArr[i] > rightArr[i] {
			totalDiff += leftArr[i] - rightArr[i]
		} else {
			totalDiff += rightArr[i] - leftArr[i]
		}
	}
	fmt.Printf("Total diff: %d\n", totalDiff)
	return
}

func PartTwo() {
	leftArr, rightArr := getInputLists()

	numberCounts := map[int]int{}
	for _, num := range rightArr {
		numberCounts[num]++
	}

	totalDiff := 0
	for i := 0; i < len(leftArr); i++ {
		rightCount := numberCounts[leftArr[i]]
		if rightCount > 0 {
			totalDiff += leftArr[i] * rightCount
		}
	}
	fmt.Printf("Total diff: %d\n", totalDiff)
	return
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
