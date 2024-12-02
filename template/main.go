package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getInputLists() []string {
	reader := bufio.NewReader(os.Stdin)
	for true {
		text, err := reader.ReadString('\n')
		fmt.Print(text)
		if err != nil || err == io.EOF {
			break
		}
	}
	return []string{}
}

func PartOne() {

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
