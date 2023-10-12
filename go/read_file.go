package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("your_file.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line (in this example, we'll just print it)
		fmt.Println(line)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}
