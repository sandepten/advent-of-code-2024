package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalSum := 0
	readFile, err := os.Open("fifthInput.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	rules := make(map[int][]int)
	updates := [][]int{}

	updateTextStarting := false
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			updateTextStarting = true
			continue
		}
		if updateTextStarting {
			updateTextStrings := strings.Split(fileScanner.Text(), ",")
			updateTextInt := []int{}

			for _, v := range updateTextStrings {
				if value, err := strconv.Atoi(v); err == nil {
					updateTextInt = append(updateTextInt, value)
				}
			}
			updates = append(updates, updateTextInt)
			continue
		}

		// rules
		ruleString := strings.Split(fileScanner.Text(), "|")
		ruleInt := []int{}
		for _, v := range ruleString {
			if value, err := strconv.Atoi(v); err == nil {
				ruleInt = append(ruleInt, value)
			}
		}
		if i, ok := rules[ruleInt[0]]; ok {
			rules[ruleInt[0]] = append(i, ruleInt[1])
		} else {
			rules[ruleInt[0]] = []int{ruleInt[1]}
		}
	}

	for _, v := range updates {
		totalSum += ruleChecker(v, rules)
	}
	fmt.Println("part1: ", totalSum)

	totalSum = 0
	for _, v := range updates {
		totalSum += ruleCheckerFixer(v, rules)
	}
	fmt.Println("part2: ", totalSum)

	readFile.Close()
}

func ruleChecker(print []int, rules map[int][]int) int {
	for i, curr := range print {
		currRules, _ := rules[curr]
		for j := i + 1; j < len(print); j++ {
			if !contains(currRules, print[j]) {
				return 0
			}
		}
	}

	mid := len(print) / 2
	return print[mid]
}

func ruleCheckerFixer(print []int, rules map[int][]int) int {
	// Function to check if the print array is in correct order
	isValidOrder := func(print []int, rules map[int][]int) bool {
		for i := 0; i < len(print)-1; i++ {
			key := print[i]
			if afterRules, exists := rules[key]; exists {
				found := false
				for _, val := range afterRules {
					if val == print[i+1] {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
		}
		return true
	}

	// Fix the order based on the rules
	fixedPrint := make([]int, len(print))
	copy(fixedPrint, print)

	// Sorting the print array according to the rules
	for i := 0; i < len(fixedPrint); i++ {
		for j := i + 1; j < len(fixedPrint); j++ {
			key := fixedPrint[i]
			if afterRules, exists := rules[key]; exists {
				shouldBeAfter := false
				for _, val := range afterRules {
					if val == fixedPrint[j] {
						shouldBeAfter = true
						break
					}
				}
				if !shouldBeAfter {
					// Swap the elements to fix the order
					fixedPrint[i], fixedPrint[j] = fixedPrint[j], fixedPrint[i]
				}
			}
		}
	}

	// If no changes were made, return 0
	if isValidOrder(print, rules) {
		return 0
	}

	// Calculate the middle index after fixing
	middle := len(fixedPrint) / 2
	return fixedPrint[middle]
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
