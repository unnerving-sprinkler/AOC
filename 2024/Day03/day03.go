package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

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

func Run() {
	fmt.Println("Starting Day 3")
	starttime := time.Now()
	answer1 := 0

	// Read Input File
	lines := inputtolines("2024/Day03/puzzleinput.txt")

	// Find All Expressions That Match Format and Add a Array Of Strings
	var matchingexpressions []string
	for _, line := range lines {
		// Search For regex Expressions
		re := regexp.MustCompile(`mul\(\d+,\d+\)`) // Expression
		matchingexpressions = append(matchingexpressions, re.FindAllString(line, -1)...)
	}

	for _, expr := range matchingexpressions { // Process Each Expression
		re := regexp.MustCompile(`\d+`)
		numbers := re.FindAllString(expr, -1)
		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])
		answer1 += number1 * number2
	}

	fmt.Println("Finished 3A in", time.Since(starttime))
	fmt.Printf("Part 1 Answer: %d\n", answer1)

	/////////////////////////// Part 2 ///////////////////////////
	answer2 := 0
	starttime = time.Now()

	matchingexpressions = nil // Reset Matching Expressions
	for _, line := range lines {
		// Search For regex Expressions
		re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`) // Expression
		matchingexpressions = append(matchingexpressions, re.FindAllString(line, -1)...)
	}

	currentmode := true // true = do, false = don't
	for _, expr := range matchingexpressions {
		if expr == "do()" { // Set Mode True If We See Command
			currentmode = true
		}
		if expr == "don't()" { // Set Mode False If We See That Command
			currentmode = false
		}
		if currentmode && expr[:3] == "mul" { // Only Process Multiplications In Do Mode
			re := regexp.MustCompile(`\d+`)
			numbers := re.FindAllString(expr, -1)
			number1, _ := strconv.Atoi(numbers[0])
			number2, _ := strconv.Atoi(numbers[1])
			answer2 += number1 * number2
		}
	}

	fmt.Println("\nFinished 3B in", time.Since(starttime))
	fmt.Printf("Part 2 Answer: %d\n", answer2)

}
