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
	fmt.Println(totalSum)

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

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
