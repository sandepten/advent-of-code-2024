package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	totalCount := 0
	readFile, err := os.Open("forthInput.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	xmasMatrix := [][]string{}

	lineNumber := 0
	for fileScanner.Scan() {
		lineArray := []string{}
		for _, v := range strings.Split(fileScanner.Text(), "") {
			lineArray = append(lineArray, v)
		}
		xmasMatrix = append(xmasMatrix, lineArray)
		lineNumber++
	}
	totalCount += countHorizontal(xmasMatrix)
	totalCount += countVertical(xmasMatrix)
	totalCount += countDiagonal(xmasMatrix)
	fmt.Println(totalCount)

	readFile.Close()
}

func countHorizontal(xmasMatrix [][]string) int {
	xmasCount := 0
	for i := 0; i < len(xmasMatrix); i++ {
		for j := 0; j < len(xmasMatrix[i])-3; j++ {
			currWord := ""
			for k := j; k < j+4; k++ {
				currWord += xmasMatrix[i][k]
			}
			if currWord == "XMAS" || currWord == "SAMX" {
				xmasCount++
			}
		}
	}

	return xmasCount
}

func countVertical(xmasMatrix [][]string) int {
	xmasCount := 0

	for i := 0; i < len(xmasMatrix[0]); i++ {
		for j := 0; j < len(xmasMatrix)-3; j++ {
			currWord := ""
			for k := j; k < j+4; k++ {
				currWord += xmasMatrix[k][i]
			}
			if currWord == "XMAS" || currWord == "SAMX" {
				xmasCount++
			}
		}
	}

	return xmasCount
}

func countDiagonal(xmasMatrix [][]string) int {
	xmasCount := 0

	for i := 0; i < len(xmasMatrix)-3; i++ {
		for j := 0; j < len(xmasMatrix[i]); j++ {
			currWord := ""

			// left Diagonal
			if j > 2 {
				for k := 0; k < 4; k++ {
					currWord += xmasMatrix[i+k][j-k]
				}
			}
			if currWord == "XMAS" || currWord == "SAMX" {
				xmasCount++
			}

			// // right Diagonal
			currWord = ""
			if j < len(xmasMatrix[i])-3 {
				for k := 0; k < 4; k++ {
					currWord += xmasMatrix[i+k][j+k]
				}
			}
			if currWord == "XMAS" || currWord == "SAMX" {
				xmasCount++
			}
		}
	}

	return xmasCount
}
