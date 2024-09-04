package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// worker simulates a part of an assembly process. It logs the start and completion of its work,
// sleeps for a random duration to mimic processing time, and signals the wait group upon completion.
func worker(part string) {
	log.Println(part, "worker begins part")     // Log the start of the work.
	time.Sleep(time.Duration(rand.Int63n(1e6))) // Simulate work by sleeping.
	log.Println(part, "worker completes part")  // Log the completion of the work.
	wg.Done()                                   // Signal the wait group that this worker has completed its work.
}

var (
	partList    = []string{"A", "B", "C", "D"} // List of parts to be processed in the assembly.
	nAssemblies = 3                            // Number of assembly cycles to be completed.
	wg          sync.WaitGroup                 // Wait group to synchronize the completion of all parts in each assembly cycle.
)

// main orchestrates the assembly process across multiple cycles. It seeds the random number generator,
// iterates through assembly cycles, and within each cycle, it launches worker goroutines for each part.
// It waits for all parts to be processed before logging the completion of the cycle.
func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
	for c := 1; c <= nAssemblies; c++ {
		log.Println("begin assembly cycle", c) // Log the start of an assembly cycle.
		wg.Add(len(partList))                  // Prepare the wait group for the number of parts.
		for _, part := range partList {
			go worker(part) // Launch a goroutine for each part.
		}
		wg.Wait()                                      // Wait for all parts to be processed.
		log.Println("assemble.  cycle", c, "complete") // Log the completion of the assembly cycle.
	}
}
