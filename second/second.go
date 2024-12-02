package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalSafe := 0
	readFile, err := os.Open("secondInput.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		list := []int{}
		parts := strings.Split(fileScanner.Text(), " ")
		for _, v := range parts {
			if value, err := strconv.Atoi(v); err == nil {
				list = append(list, value)
			}
		}
		if !checkListSafetiness(list) {
			for i := 0; i < len(list); i++ {
				listProblemDampner := removeElement(list, i)
				if checkListSafetiness(listProblemDampner) {
					totalSafe++
					break
				}
			}
		} else {
			totalSafe++
		}
	}
	fmt.Println(totalSafe)

	readFile.Close()
}

func checkListSafetiness(list []int) bool {
	isListIncreasing := list[0] < list[1]
	for i := 1; i < len(list); i++ {
		difference := int(math.Abs(float64(list[i-1] - list[i])))
		if difference < 1 || difference > 3 {
			return false
		}
		if (isListIncreasing && list[i-1] > list[i]) || (!isListIncreasing && list[i-1] < list[i]) {
			return false
		}
	}
	return true
}

func removeElement(list []int, index int) []int {
	result := make([]int, 0, len(list)-1)
	result = append(result, list[:index]...)
	result = append(result, list[index+1:]...)
	return result
}
