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
	fmt.Println(totalCount)

	readFile.Close()
}

func countHorizontal(xmasMatrix [][]string) {

}
