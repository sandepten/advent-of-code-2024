package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	readFile, err := os.Open("seventhInput.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	calibrations := make(map[int][]int)

	for fileScanner.Scan() {
		lineSplit := strings.Split(fileScanner.Text(), ": ")
		combinations := []int{}
		for _, v := range strings.Split(lineSplit[1], " ") {
			numValue, _ := strconv.Atoi(v)
			combinations = append(combinations, numValue)
		}
		key, _ := strconv.Atoi(lineSplit[0])
		calibrations[key] = combinations
	}

	total = part1Ans(calibrations)
	fmt.Println(total)

	readFile.Close()
}

func part1Ans(calibrations map[int][]int) int {
	totalSum := 0
	for key, value := range calibrations {
		if isCalibrationCorrect(key, value) {
			totalSum += key
		}
	}
	return totalSum
}

/*
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
Operators are always evaluated left-to-right, not according to precedence rules. Furthermore, numbers in the equations cannot be rearranged. Glancing into the jungle, you can see elephants holding two different types of operators: add (+) and multiply (*).

Only three of the above equations can be made true by inserting operators:

190: 10 19 has only one position that accepts an operator: between 10 and 19. Choosing + would give 29, but choosing * would give the test value (10 * 19 = 190).
3267: 81 40 27 has two positions for operators. Of the four possible configurations of the operators, two cause the right side to match the test value: 81 + 40 * 27 and 81 * 40 + 27 both equal 3267 (when evaluated left-to-right)!
292: 11 6 16 20 can be solved in exactly one way: 11 + 6 * 16 + 20.

here key is the result of the equation and value is the numbers in the equation
*/
func isCalibrationCorrect(key int, value []int) bool {
	// Try all possible operator configurations
	for i := 0; i < (1 << (len(value) - 1)); i++ {
		// Evaluate the expression with the current operator configuration
		if evaluateExpression(value, i) == key {
			return true
		}
	}
	return false
}

func evaluateExpression(nums []int, operatorConfig int) int {
	result := nums[0]
	for j := 0; j < len(nums)-1; j++ {
		// Check if the jth bit is set in operatorConfig
		// If bit is 0, use addition (+)
		// If bit is 1, use multiplication (*)
		if (operatorConfig & (1 << j)) == 0 {
			// Addition
			result += nums[j+1]
		} else {
			// Multiplication
			result *= nums[j+1]
		}
	}
	return result
}
