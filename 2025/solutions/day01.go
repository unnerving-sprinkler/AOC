package solutions

import (
	"fmt"
	"strconv"
	"time"
)

func Run() {

	// Setup Part 1
	starttime := time.Now()
	answer1 := 0
	testtag := 0
	dialpos := 50
	fmt.Println("Starting 2025 Day 1 Part 1")

	// Setup VARS
	type Momement struct {
		Direction string
		Distance  int
	}
	instructions := []Momement{}

	// Read Input
	lines := inputtolines("2025/puzzleinput/day01.txt")

	// Put Input Into structure
	for _, line := range lines {
		dir := line[0:1]
		dist, _ := strconv.Atoi(line[1:])

		m := Momement{
			Direction: dir,
			Distance:  dist,
		}
		instructions = append(instructions, m)
	}

	// Process Each Instruction
	for _, inst := range instructions {
		if inst.Direction == "L" {
			dialpos -= inst.Distance
		} else if inst.Direction == "R" {
			dialpos += inst.Distance
		}

		// Wrap Around

		for dialpos > 99 {
			dialpos -= 100

		}

		for dialpos < 0 {
			dialpos += 100

		}

		if dialpos == 0 {
			answer1++
		}
	}

	//Close Part 1
	fmt.Println("Finished Part 1 in", time.Since(starttime))
	fmt.Printf("Part 1 Answer: %d\n", answer1)

	fmt.Println(testtag)
	////////////////////////////// Part 2 \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

	// Setup Part 2
	starttime = time.Now()
	answer2 := 0
	fmt.Println("\nStarting 2025 Day 1 Part 2")

	// Close Part 2
	fmt.Println("Finished Part 2 in", time.Since(starttime))
	fmt.Printf("Part 2 Answer: %d\n", answer2)

}
