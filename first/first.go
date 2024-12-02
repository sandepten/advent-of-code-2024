package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	list1, list2 := getLists()

	// ! Part 1
	// sort the lists
	slices.Sort(list1)
	slices.Sort(list2)

	fmt.Println("Part 1: ", differenceCalculator(list1, list2))

	//! Part 2
	fmt.Println("Part 2: ", totalInstancesMultiplier(list1, list2))
}

func getLists() ([]int, []int) {
	list1 := []int{}
	list2 := []int{}

	readFile, err := os.Open("firstInput.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), "   ")
		if value, err := strconv.Atoi(parts[0]); err == nil {
			list1 = append(list1, value)
		}
		if value, err := strconv.Atoi(parts[1]); err == nil {
			list2 = append(list2, value)
		}
	}

	readFile.Close()

	return list1, list2
}

func differenceCalculator(list1, list2 []int) int {
	totalSum := 0
	for i := 0; i < len(list1); i++ {
		totalSum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return totalSum
}

func totalInstancesMultiplier(list1, list2 []int) int {
	totalSum := 0
	list2Instances := make(map[int]int)
	for _, v := range list2 {
		if i, ok := list2Instances[v]; ok {
			list2Instances[v] = i + 1
		} else {
			list2Instances[v] = 1
		}
	}
	for _, v := range list1 {
		if i, ok := list2Instances[v]; ok {
			totalSum += v * i
		}
	}
	return totalSum
}
