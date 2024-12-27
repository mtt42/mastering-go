package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments.")
		return
	}

	var total, nInt, nFloat int
	invalid := make([]string, 0)

	for _, k := range arguments[1:] {
		// Is it and integer?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInt++
			continue
		}
		// Is it a float?
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloat++
			continue
		}

		invalid = append(invalid, k)

	}
	fmt.Printf("Total: %d, Integers: %d, Floats: %d\n", total, nInt, nFloat)
	if len(invalid) > 0 {
		fmt.Printf("Invalid arguments: %v\n", invalid)
	}
}
