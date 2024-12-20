package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Constants representing the types of items needed to smoke.
const (
	paper = iota // Represents paper.
	grass        // Represents grass.
	match        // Represents a match.
)

// smokeMap maps the item constants to their string representations.
var smokeMap = map[int]string{
	paper: "paper",
	grass: "grass",
	match: "match",
}

// names maps the item constants to the names of the smokers who need them.
var names = map[int]string{
	paper: "Sandy",
	grass: "Apple",
	match: "Daisy",
}

// Table represents the table with channels for each item.
type Table struct {
	paper chan int // Channel for paper.
	grass chan int // Channel for grass.
	match chan int // Channel for match.
}

// arbitrate function simulates the arbitrator's role in the cigarette smokers problem.
// It randomly selects an item to place on the table and notifies the smokers.
func arbitrate(t *Table, smokers [3]chan int) {
	for {
		time.Sleep(time.Millisecond * 500) // Simulate time between selections.
		//nolint:all
		next := rand.Intn(3) // Randomly select an item to place on the table.
		fmt.Printf("Table chooses %s: %s\n", smokeMap[next], names[next])
		switch next {
		case paper:
			t.grass <- 1 // Place grass on the table.
			t.match <- 1 // Place a match on the table.
		case grass:
			t.paper <- 1 // Place paper on the table.
			t.match <- 1 // Place a match on the table.
		case match:
			t.grass <- 1 // Place grass on the table.
			t.paper <- 1 // Place paper on the table.
		}
		for _, smoker := range smokers {
			smoker <- next // Notify all smokers of the item placed on the table.
		}
		wg.Add(1)
		wg.Wait() // Wait for the smoker to finish smoking.
	}
}

// smoker function simulates a smoker who waits for their required items to be placed on the table.
// Once the items are available, they "smoke" and then signal completion.
func smoker(t *Table, name string, smokes int, signal chan int) {
	var chosen int
	for {
		chosen = <-signal // Wait for notification of an item placed on the table.

		if smokes != chosen {
			continue // If the item is not the one needed, continue waiting.
		}

		// Simulate taking the items from the table.
		fmt.Printf("Table: %d grass: %d match: %d\n", len(t.paper), len(t.grass), len(t.match))
		select {
		case <-t.paper:
		case <-t.grass:
		case <-t.match:
		}
		// Simulate smoking process.
		fmt.Printf("%s smokes a cigarette\n", name)
		time.Sleep(time.Millisecond * 500) // Simulate time taken to smoke.
		wg.Done()                          // Signal that smoking is complete.
		time.Sleep(time.Millisecond * 100) // Simulate waiting before ready to smoke again.
	}
}

const LIMIT = 1 // Limit for the channels to simulate the table's capacity.

var wg *sync.WaitGroup // WaitGroup to synchronize smokers' smoking process.

// main function sets up the simulation of the cigarette smokers problem.
// It initializes the table, the smokers, and starts the arbitration process.
func main() {
	wg = new(sync.WaitGroup)
	table := new(Table)
	// Initialize channels with a buffer to simulate the table's capacity.
	table.match = make(chan int, LIMIT)
	table.paper = make(chan int, LIMIT)
	table.grass = make(chan int, LIMIT)
	var signals [3]chan int
	// Initialize and start goroutines for each smoker.
	for i := 0; i < 3; i++ {
		signal := make(chan int, 1)
		signals[i] = signal
		go smoker(table, names[i], i, signal)
	}
	// Start the arbitration process.
	arbitrate(table, signals)
}
