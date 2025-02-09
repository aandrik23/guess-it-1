package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Stores the sequence of numbers read from input
var numbers []int

// Calculates the average of a slice of integers
func Average(arr []int) float64 {
	total := 0
	for _, num := range arr {
		total += num
	}
	return float64(total) / float64(len(arr))
}

// Determines the predicted range for the next number
func predictRange(current, previous int) (int, int) {
	// Compute range as ±10% of the difference between current and previous number
	delta := float64(current - previous)
	lower := int(float64(current) - delta*0.1)
	upper := int(float64(current) + delta*0.1)
	return lower, upper
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read numbers from standard input
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Skipping invalid number %q: %v", line, err)
			continue
		}

		// Append the number to the list
		numbers = append(numbers, num)

		// Ensure at least two numbers exist before predicting a range
		if len(numbers) > 1 {
			previous := numbers[len(numbers)-2]
			lower, upper := predictRange(num, previous)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}

	// Check for errors during input reading
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
