package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getInputLists() ([][]int, [][]int) {
	reader := bufio.NewReader(os.Stdin)
	rules := [][]int{}
	updates := [][]int{}
	processingRules := true
	for true {
		text, err := reader.ReadString('\n')
		if text == "\n" {
			processingRules = false
			continue
		} else if text == "" {
			break
		}

		if processingRules {
			pages := strings.Split(strings.TrimRight(text, "\n"), "|")
			first, err := strconv.Atoi(pages[0])
			if err != nil {
				log.Fatalf("Failed to parse int from text: %s\n", pages[0])
			}

			second, err := strconv.Atoi(pages[1])
			if err != nil {
				log.Fatalf("Failed to parse int from text: %s\n", pages[1])
			}
			rules = append(rules, []int{first, second})
		} else {
			pages := strings.Split(strings.TrimRight(text, "\n"), ",")
			update := []int{}
			for _, page := range pages {
				val, err := strconv.Atoi(page)
				if err != nil {
					log.Fatalf("Failed to parse update page: %s\n", page)
				}
				update = append(update, val)
			}
			updates = append(updates, update)
		}

		if err != nil || err == io.EOF {
			break
		}
	}
	return rules, updates
}

func PartOne() {
	// Rules is a list of pairs of page numbers
	// The first item must come before the second item
	rules, updates := getInputLists()
	fmt.Printf("Rules:\n%v\n", rules)
	fmt.Printf("Updates:\n%v\n", updates)

	cantFollowRuleMap := make(map[int][]int)
	for _, newRule := range rules {
		rule, _ := cantFollowRuleMap[newRule[1]]
		cantFollowRuleMap[newRule[1]] = append(rule, newRule[0])
	}

	total := 0
	for _, update := range updates {
		invalidPage := false
		curPages := []int{}
		badNewPages := []int{}
		for _, page := range update {
			// Check if new page is allowed
			if slices.Contains(badNewPages, page) {
				invalidPage = true
				break
			}

			// Grab additional bad pages (pages that can't follow this page)
			rules, _ := cantFollowRuleMap[page]
			for _, page := range rules {
				badNewPages = append(badNewPages, page)
			}
			curPages = append(curPages, page)
		}

		// Add middle page to total
		if !invalidPage {
			middlePage := update[len(update)/2]
			total += middlePage
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func PartTwo() {
	// Rules is a list of pairs of page numbers
	// The first item must come before the second item
	rules, updates := getInputLists()
	fmt.Printf("Rules:\n%v\n", rules)
	fmt.Printf("Updates:\n%v\n", updates)

	cantFollowRuleMap := make(map[int][]int)
	for _, newRule := range rules {
		rule, _ := cantFollowRuleMap[newRule[1]]
		cantFollowRuleMap[newRule[1]] = append(rule, newRule[0])
	}

	badUpdates := []int{}
	for i, update := range updates {
		invalidPage := false
		curPages := []int{}
		badNewPages := []int{}
		for _, page := range update {
			// Check if new page is allowed
			if slices.Contains(badNewPages, page) {
				invalidPage = true
				break
			}

			// Grab additional bad pages (pages that can't follow this page)
			rules, _ := cantFollowRuleMap[page]
			for _, page := range rules {
				badNewPages = append(badNewPages, page)
			}
			curPages = append(curPages, page)
		}

		// Add to list for next processing step
		if invalidPage {
			badUpdates = append(badUpdates, i)
		}
	}

	// Fix updates by sorting using ruleset as comparison func
	total := 0
	for _, updateIdx := range badUpdates {
		update := updates[updateIdx]
		// if a in b's cantfollowmap, then a must come after b
		slices.SortFunc(update, func(a, b int) int {
			aCantFollow, _ := cantFollowRuleMap[a]
			bCantFollow, _ := cantFollowRuleMap[b]
			if slices.Contains(aCantFollow, b) {
				return 1
			} else if slices.Contains(bCantFollow, a) {
				return -1
			} else {
				return 0
			}
		})
		total += update[len(update)/2]
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
