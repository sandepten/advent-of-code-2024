package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	totalSteps := 0
	readFile, err := os.Open("sixthInput.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	guardPosition := []int{0, 0}
	guardMap := [][]string{}

	lineNumber := 0
	for fileScanner.Scan() {
		lineArray := []string{}
		for index, v := range strings.Split(fileScanner.Text(), "") {
			if v == "^" {
				guardPosition[0] = index
				guardPosition[1] = lineNumber
			}
			lineArray = append(lineArray, v)
		}
		guardMap = append(guardMap, lineArray)
		lineNumber++
	}

	totalSteps = reachOutside(guardMap, guardPosition)

	fmt.Println(totalSteps)

	readFile.Close()
}

func reachOutside(guardMap [][]string, guardPosition []int) int {
	totalSteps := 1
	guardPositionX := guardPosition[0]
	guardPositionY := guardPosition[1]
	currDirection := "up"

	distinctMatrix := [][]bool{}
	for i := 0; i < len(guardMap); i++ {
		distinctLine := make([]bool, len(guardMap[i]))
		distinctMatrix = append(distinctMatrix, distinctLine)
	}

	for guardPositionX > 0 && guardPositionX < len(guardMap[0]) && guardPositionY > 0 && guardPositionY < len(guardMap) {
		if currDirection == "up" {
			if guardPositionY == 0 {
				break
			}
			if guardMap[guardPositionY-1][guardPositionX] == "#" {
				currDirection = "right"
				continue
			}
			if distinctMatrix[guardPositionY][guardPositionX] == true {
				guardPositionY--
				continue
			}
			distinctMatrix[guardPositionY][guardPositionX] = true
			guardPositionY--
		} else if currDirection == "down" {
			if guardPositionY == len(guardMap)-1 {
				break
			}
			if guardMap[guardPositionY+1][guardPositionX] == "#" {
				currDirection = "left"
				continue
			}
			if distinctMatrix[guardPositionY][guardPositionX] == true {
				guardPositionY++
				continue
			}
			distinctMatrix[guardPositionY][guardPositionX] = true
			guardPositionY++
		} else if currDirection == "right" {
			if guardPositionX == len(guardMap[0])-1 {
				break
			}
			if guardMap[guardPositionY][guardPositionX+1] == "#" {
				currDirection = "down"
				continue
			}
			if distinctMatrix[guardPositionY][guardPositionX] == true {
				guardPositionX++
				continue
			}
			distinctMatrix[guardPositionY][guardPositionX] = true
			guardPositionX++
		} else if currDirection == "left" {
			if guardPositionX == 0 {
				break
			}
			if guardMap[guardPositionY][guardPositionX-1] == "#" {
				currDirection = "up"
				continue
			}
			if distinctMatrix[guardPositionY][guardPositionX] == true {
				guardPositionX--
				continue
			}
			distinctMatrix[guardPositionY][guardPositionX] = true
			guardPositionX--
		}
		totalSteps++
	}

	return totalSteps
}
