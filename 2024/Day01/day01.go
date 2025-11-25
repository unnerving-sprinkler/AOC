package day01

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

func Run() {
	fmt.Println("Starting Day 1")
	starttime := time.Now()

	// Open File Read To Stringgs
	path := "Day01/puzzleinput.txt"
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
	var answer1 float64
	for n := 0; n < len(firstnumbers); n++ {
		answer1 += math.Abs(float64(firstnumbers[n]) - float64(secondnumbers[n]))
	}

	// Report Results
	fmt.Println("Finished Part 1 in", time.Now().Sub(starttime))
	fmt.Printf("Part 1 Answer: %.0f\n", answer1)

	// Part 2
	var answer2 int
	for _, firstnumber := range firstnumbers {
		n := 0
		for _, secondnumber := range secondnumbers {
			if firstnumber == secondnumber {
				n++
			}

		}
		answer2 += firstnumber * n
	}

	// Report Results
	fmt.Println("\nFinished Part 2 in", time.Since(starttime))
	fmt.Printf("Part 2 Answer: %d\n", answer2)

}
