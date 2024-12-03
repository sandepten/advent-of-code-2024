package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	totalSum := 0
	readFile, err := os.Open("thirdInput.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	shouldMultipleDoCondition := true
	for fileScanner.Scan() {
		totalSum += totalSumForLine(fileScanner.Text(), &shouldMultipleDoCondition)
	}
	fmt.Println(totalSum)

	readFile.Close()
}

func totalSumForLine(line string, shouldMultipleDoCondition *bool) int {
	totalSum := 0
	mulRegex := `^mul\((-?\d+),(-?\d+)\)$`
	regex := regexp.MustCompile(mulRegex)

	subStrings := getConsecutiveSubstrings(line)
	for _, v := range subStrings {
		if v == "don't()" {
			*shouldMultipleDoCondition = false
		}
		if v == "do()" {
			*shouldMultipleDoCondition = true
		}
		if !*shouldMultipleDoCondition {
			continue
		}
		if matches := regex.FindStringSubmatch(v); matches != nil {
			// Extract the numbers
			num1, _ := strconv.Atoi(matches[1])
			num2, _ := strconv.Atoi(matches[2])

			totalSum += num1 * num2
		}
	}

	return totalSum
}

func getConsecutiveSubstrings(s string) []string {
	// If the string is empty, return an empty slice
	if len(s) == 0 {
		return []string{}
	}

	substrings := []string{}

	for start := 0; start < len(s); start++ {
		for end := start + 1; end <= len(s); end++ {
			substrings = append(substrings, s[start:end])
		}
	}

	return substrings
}
