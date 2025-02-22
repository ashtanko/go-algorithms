package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	sleeping = iota
	checking
	cutting
)

var stateLog = map[int]string{
	0: "Sleeping",
	1: "Checking",
	2: "Cutting",
}
var wg *sync.WaitGroup // Amount of potential customers

type Barber struct {
	name     string
	mutex    sync.Mutex
	state    int // Sleeping/Checking/Cutting
	customer *Customer
}

type Customer struct {
	name string
}

func (c *Customer) String() string {
	return fmt.Sprintf("%p", c)[7:]
}

func NewBarber() (b *Barber) {
	return &Barber{
		name:  "Sam",
		state: sleeping,
	}
}

// Barber goroutine
// Checks for customers
// Sleeps - wait for wakes to wake him up
func barber(b *Barber, wr chan *Customer, wakes chan *Customer) {
	for {
		b.mutex.Lock()
		// defer b.Unlock()
		b.state = checking
		b.customer = nil

		// checking the waiting room
		fmt.Printf("Checking waiting room: %d\n", len(wr))
		time.Sleep(time.Millisecond * 100)
		select {
		case c := <-wr:
			HairCut(c, b)
			b.mutex.Unlock()
		default: // Waiting room is empty
			fmt.Printf("Sleeping Barber - %s\n", b.customer)
			b.state = sleeping
			b.customer = nil
			b.mutex.Unlock()
			c := <-wakes
			b.mutex.Lock()
			fmt.Printf("Woken by %s\n", c)
			HairCut(c, b)
			b.mutex.Unlock()
		}
	}
}

func HairCut(c *Customer, b *Barber) {
	b.state = cutting
	b.customer = c
	b.mutex.Unlock()
	fmt.Printf("Cutting  %s hair\n", c)
	time.Sleep(time.Millisecond * 100)
	b.mutex.Lock()
	wg.Done()
	b.customer = nil
}

// customer goroutine
// just fizzles out if it's full, otherwise the customer
// is passed along to the channel handling it's haircut etc
func customer(c *Customer, b *Barber, wr chan<- *Customer, wakers chan<- *Customer) {
	// arrive
	time.Sleep(time.Millisecond * 50)
	// Check on barber
	b.mutex.Lock()
	fmt.Printf("Customer %s checks %s barber | room: %d, w %d - customer: %s\n",
		c, stateLog[b.state], len(wr), len(wakers), b.customer)
	switch b.state {
	case sleeping:
		select {
		case wakers <- c:
		default:
			select {
			case wr <- c:
			default:
				wg.Done()
			}
		}
	case cutting:
		select {
		case wr <- c:
		default: // Full waiting room, leave shop
			wg.Done()
		}
	case checking:
		panic("Customer shouldn't check for the Barber when Barber is Checking the waiting room")
	}
	b.mutex.Unlock()
}

func main() {
	b := NewBarber()
	b.name = "Rocky"
	WaitingRoom := make(chan *Customer, 5) // 5 chairs
	Wakes := make(chan *Customer, 1)       // Only one wake at a time
	go barber(b, WaitingRoom, Wakes)

	time.Sleep(time.Millisecond * 100)
	wg = new(sync.WaitGroup)
	n := 10
	wg.Add(10)
	// Spawn customers
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 50)
		c := &Customer{
			name: fmt.Sprintf("customer: %d", i),
		}
		go customer(c, b, WaitingRoom, Wakes)
	}

	wg.Wait()
	fmt.Println("No more customers for the day")
}
