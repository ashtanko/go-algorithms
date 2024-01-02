package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each
// pair of adjacent philosophers. Each philosopher must alternately think and eat. However, a philosopher
// can only eat spaghetti when they have both left and right forks. Each fork can be held by only one
// philosopher and so a philosopher can use the fork only if it is not being used by another philosopher.
// After an individual philosopher finishes eating, they need to put down both forks so that the forks
// become available to others. A philosopher can take the fork on their right or the one on their left as
// they become available, but cannot start eating before getting both forks. Eating is not limited by the
// remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed.
// The problem is how to design a discipline of behavior (a concurrent algorithm) such that no philosopher
// will starve; i.e., each can forever continue to alternate between eating and thinking, assuming that no
// philosopher can know when others may want to eat or think.

// Number of philosophers is simply the length of this list.
var ph = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}

const hunger = 3                // Number of times each philosopher eats
const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time

var fmt = log.New(os.Stdout, "", 0)

var dining sync.WaitGroup

func diningProblem(phName string, dominantHand, otherHand *sync.Mutex) {
	fmt.Println(phName, "Seated")
	h := fnv.New64a()
	_, err := h.Write([]byte(phName))
	if err != nil {
		fmt.Printf("Error to write %s", phName)
		return
	}
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		fmt.Println(phName, "Hungry")
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		fmt.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		fmt.Println(phName, "Thinking")
		rSleep(think)
	}
	fmt.Println(phName, "Satisfied")
	dining.Done()
	fmt.Println(phName, "Left the table")
}

func main() {
	fmt.Println("Table empty")
	dining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph[0], fork0, forkLeft)
	dining.Wait() // wait for philosophers to finish
	fmt.Println("Table empty")
}
