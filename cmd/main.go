package main

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
)

func main() {
	// Example of using the MeasureMemoryUsage function to measure memory usage
	memoryUsage := utils.MeasureMemoryUsage(func() {
		// The code you want to measure goes here
		for i := 0; i < 1000000; i++ {
			_ = make([]byte, 1024) // Allocating memory for the example
		}
	})

	fmt.Printf("Memory used: %d bytes\n", memoryUsage)
}
