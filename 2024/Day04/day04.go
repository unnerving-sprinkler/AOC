package day04

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func inputtolines(path string) []string { // Read Input File Into Lines

	var lines []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open file:", err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines
}

func savegrid(grid [][]string) []string {
	var lines []string
	for _, row := range grid {
		lines = append(lines, strings.Join(row, ""))
	}
	return lines
}

func countxmas(lines []string) int {
	var matchingexpressions []string
	re := regexp.MustCompile(`XMAS|SAMX`) // Expression
	for _, line := range lines {
		matchingexpressions = append(matchingexpressions, re.FindAllString(line, -1)...)
	}
	return len(matchingexpressions)
}

func Run() {
	fmt.Println("Starting Day 4")
	starttime := time.Now()
	answer1 := 0

	var grid [][]string
	// Read Input File
	lines := inputtolines("Day04/puzzleinput.txt")

	for _, line := range lines { // Create a 2D Array From Input
		grid = append(grid, strings.Split(line, ""))
	}

	normaldirection := savegrid(grid)

	fmt.Println(normaldirection)

	fmt.Println("Finished 4A in", time.Since(starttime))
	fmt.Printf("Part 1 Answer: %d\n", answer1)

	/////////////////////////// Part 2 ///////////////////////////
	answer2 := 0
	starttime = time.Now()

	fmt.Println("\nFinished 4B in", time.Since(starttime))
	fmt.Printf("Part 2 Answer: %d\n", answer2)

}
