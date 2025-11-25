package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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

func stringstonumbers(strings []string) []int { //Convert Array of Strings to Array of Ints
	var numbers []int
	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		numbers = append(numbers, num)
	}
	return numbers
}

func reverseInts(numbers []int) []int { // Reverse Array of Ints
	reversed := make([]int, len(numbers))
	for i, num := range numbers {
		reversed[len(numbers)-1-i] = num
	}
	return reversed
}

func issafereport(numbers []int) bool { // Review Report Using AOC Rules
	// Check If All Assending Or Decending
	issorted := sort.IntsAreSorted(numbers) || sort.IntsAreSorted(reverseInts(numbers))

	// Check Distance Each Number Is from Last
	rate := true
	for n, _ := range numbers[1:] {
		diff := math.Abs(float64(numbers[n+1] - numbers[n]))
		if diff < 1 || diff > 3 {
			rate = false
		}
	}
	if issorted && rate {
		return true
	} else {
		return false
	}
}

func main() {
	// Starting Day 2
	fmt.Println("Starting Day 2")
	answer1 := 0
	starttime := time.Now()

	// Read Input Into Lines
	path := "puzzleinput.txt"
	lines := inputtolines(path)

	// Process Line One at a Time
	for _, line := range lines {
		strings := strings.Split(line, " ") // Split Line Into Array
		numbers := stringstonumbers(strings)

		if issafereport(numbers) {
			answer1 += 1
		}

	}

	// Output Answer
	fmt.Println("Finished Part 1 in", time.Since(starttime))
	fmt.Printf("Part 1 Answer: %d\n", answer1)

}
