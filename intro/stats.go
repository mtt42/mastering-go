package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// The main function is the entry point of the program.
func main() {
	// os.Args contains the command-line arguments passed to the program.
	// The first element (os.Args[0]) is the program name, so we need to start from index 1.
	arguments := os.Args

	// Check if any arguments are provided.
	// If only the program name is present (len(arguments) == 1), prompt the user for input and exit.
	if len(arguments) == 1 {
		fmt.Println("Please give one or more arguments.")
		return
	}

	// Declare variables to keep track of minimum and maximum values.
	// Since these are float64, they will hold decimal numbers.
	var min, max float64

	// Use a flag (initialized) to determine if the min/max variables are set.
	// Start with 0 to indicate they haven't been initialized yet.
	initialized := 0

	// nValues counts the number of valid float values provided by the user.
	nValues := 0

	// sum accumulates the sum of valid float values.
	var sum float64

	// Iterate over all command-line arguments, starting from the second element (index 1).
	for i := 1; i < len(arguments); i++ {
		// Convert the current argument from string to float64.
		n, err := strconv.ParseFloat(arguments[i], 64)
		// If an error occurs (e.g., the argument is not a valid number), skip this iteration.
		if err != nil {
			continue
		}

		// Increment the count of valid values.
		nValues++

		// Add the current number to the sum.
		sum += n

		// If min/max values are not initialized, set both to the current number.
		if initialized == 0 {
			min = n
			max = n
			initialized = 1 // Mark the min/max as initialized.
			continue
		}

		// Update the min value if the current number is smaller.
		if n < min {
			min = n
		}

		// Update the max value if the current number is larger.
		if n > max {
			max = n
		}
	}

	// Print the results.
	fmt.Printf("Number of values: %d\n", nValues)
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Min: %.2f\n", min)
	fmt.Printf("Max: %.2f\n", max)

	// Mean value
	if nValues == 0 {
		return
	}

	meanValue := sum / float64(nValues)
	fmt.Printf("Mean value: %.5f\n", meanValue)

	// Standard deviation
	var squared float64

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		squared = squared + math.Pow(n-meanValue, 2)
	}

	standardDeviation := math.Sqrt(squared / float64(nValues))
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)
}
