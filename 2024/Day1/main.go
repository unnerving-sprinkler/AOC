package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

func inputtolines(path string) []string {

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

func linetoints(line string) (int, int) {

	a, _ := strconv.Atoi(line[0:5])
	b, _ := strconv.Atoi(line[8:13])
	return a, b

}

func main() {
	fmt.Println("Starting Day 1")
	starttime := time.Now()

	// Open File Read To Stringgs
	path := "puzzleinput.txt"
	lines := inputtolines(path)

	// Create a Place To Store int values and seperate numbers from the string
	var firstnumbers []int
	var secondnumbers []int
	for _, line := range lines {
		l1, l2 := linetoints(line)
		firstnumbers = append(firstnumbers, l1)
		secondnumbers = append(secondnumbers, l2)
	}

	// Order Lists
	sort.Ints(firstnumbers)
	sort.Ints(secondnumbers)

	//Determine Answer
	var answer float64
	for n := 0; n < len(firstnumbers); n++ {
		answer += math.Abs(float64(firstnumbers[n]) - float64(secondnumbers[n]))
	}

	// Report Results
	finishtime := time.Now()
	fmt.Println("Finished in", finishtime.Sub(starttime))
	fmt.Println("Answer:", answer)
}
