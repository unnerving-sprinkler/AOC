package solutions

import (
	"bufio"
	"fmt"
	"os"
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
