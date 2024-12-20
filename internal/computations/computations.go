package main

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	data := strings.Repeat(strings.Repeat("a", 1000)+"\n", 1000)
	run_line_short(data, true)
	run_line_long(data, true)
	run_line_short_workers(data, true)
	run_line_long_workers(data, true)
	run_bulk_short(data, true)
	run_bulk_long(data, true)
	run_seq_short(data, true)
	run_seq_long(data, true)
}

func run_line_short(data string, stat bool) {
	if stat {
		s := stats("run_line_short")
		defer s()
	}
	r := strings.NewReader(data)
	err := process(r, line_handler_short)
	if err != nil {
		panic(err)
	}
}
func run_line_long(data string, stat bool) {
	if stat {
		s := stats("run_line_long")
		defer s()
	}
	r := strings.NewReader(data)
	err := process(r, line_handler_long)
	if err != nil {
		panic(err)
	}
}
func run_line_short_workers(data string, stat bool) {
	if stat {
		s := stats("run_line_short_workers")
		defer s()
	}
	r := strings.NewReader(data)
	err := processWorkers(r, line_handler_short)
	if err != nil {
		panic(err)
	}
}
func run_line_long_workers(data string, stat bool) {
	if stat {
		s := stats("run_line_long_workers")
		defer s()
	}
	r := strings.NewReader(data)
	err := processWorkers(r, line_handler_long)
	if err != nil {
		panic(err)
	}
}
func run_bulk_short(data string, stat bool) {
	if stat {
		s := stats("run_bulk_short")
		defer s()
	}
	r := strings.NewReader(data)
	err := processBulk(r, bulk_handler_short)
	if err != nil {
		panic(err)
	}
}
func run_bulk_long(data string, stat bool) {
	if stat {
		s := stats("run_bulk_long")
		defer s()
	}
	r := strings.NewReader(data)
	err := processBulk(r, bulk_handler_long)
	if err != nil {
		panic(err)
	}
}
func run_seq_short(data string, stat bool) {
	if stat {
		s := stats("run_seq_short")
		defer s()
	}
	r := strings.NewReader(data)
	err := processSeq(r, line_handler_short)
	if err != nil {
		panic(err)
	}
}
func run_seq_long(data string, stat bool) {
	if stat {
		s := stats("run_seq_long")
		defer s()
	}
	r := strings.NewReader(data)
	err := processSeq(r, line_handler_long)
	if err != nil {
		panic(err)
	}
}

func line_handler_short(k string) error {
	_ = len(k)
	return nil
}

func line_handler_long(k string) error {
	<-time.After(time.Millisecond * 5)
	_ = len(k)
	return nil
}

func bulk_handler_short(b []string) error {
	for _, k := range b {
		_ = len(k)
	}
	return nil
}

func bulk_handler_long(b []string) error {
	<-time.After(time.Millisecond * 5)
	for _, k := range b {
		_ = len(k)
	}
	return nil
}

func stats(name string) func() {
	fmt.Printf("======================\n")
	fmt.Printf("%v\n", name)
	start := time.Now()
	return func() {
		fmt.Printf("time to run %v\n", time.Since(start))
		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)
		fmt.Printf("Alloc: %d MB, TotalAlloc: %d MB, Sys: %d MB\n",
			ms.Alloc/1024/1024, ms.TotalAlloc/1024/1024, ms.Sys/1024/1024)
		fmt.Printf("Mallocs: %d, Frees: %d\n",
			ms.Mallocs, ms.Frees)
		fmt.Printf("HeapAlloc: %d MB, HeapSys: %d MB, HeapIdle: %d MB\n",
			ms.HeapAlloc/1024/1024, ms.HeapSys/1024/1024, ms.HeapIdle/1024/1024)
		fmt.Printf("HeapObjects: %d\n", ms.HeapObjects)
		fmt.Printf("\n")
	}
}

func process(r io.Reader, h func(string) error) error {
	errs := make(chan error)
	workers := make(chan struct{}, 4)
	var wg sync.WaitGroup
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			workers <- struct{}{} // acquire a token
			wg.Add(1)
			go func(line string) {
				defer wg.Done()
				if err := h(line); err != nil {
					errs <- err
				}
				<-workers
			}(scanner.Text())
		}
		wg.Wait()
		if err := scanner.Err(); err != nil {
			errs <- err
		}
		close(errs)
	}()
	var err error
	for e := range errs {
		if e != nil && err == nil {
			err = e
		}
	}
	return err
}

func processWorkers(r io.Reader, h func(string) error) error {
	errs := make(chan error)
	input := make(chan string)
	y := 4
	var wg sync.WaitGroup
	for i := 0; i < y; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range input {
				if err := h(line); err != nil {
					errs <- err
				}
			}
		}()
	}
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			input <- scanner.Text()
		}
		close(input)
		wg.Wait()
		if err := scanner.Err(); err != nil {
			errs <- err
		}
		close(errs)
	}()
	var err error
	for e := range errs {
		if err == nil && e != nil {
			err = e
		}
	}
	return err
}

func processBulk(r io.Reader, h func([]string) error) error {
	errs := make(chan error)
	input := make(chan []string)
	y := 4
	var wg sync.WaitGroup
	for i := 0; i < y; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for bulk := range input {
				if err := h(bulk); err != nil {
					errs <- err
				}
			}
		}()
	}
	go func() {
		scanner := bufio.NewScanner(r)
		l := 50
		bulk := make([]string, l)
		i := 0
		for scanner.Scan() {
			text := scanner.Text()
			bulk[i] = text
			i++
			if i == l {
				copied := make([]string, l)
				copy(copied, bulk)
				i = 0
				input <- copied
			}
		}
		if i > 0 {
			input <- bulk[:i]
		}
		close(input)
		if err := scanner.Err(); err != nil {
			errs <- err
		}
	}()
	go func() {
		wg.Wait()
		close(errs)
	}()
	var err error
	for e := range errs {
		if err == nil && e != nil {
			err = e
		}
	}
	return err
}

func processSeq(r io.Reader, h func(string) error) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		if err := h(text); err != nil {
			return err
		}
	}
	return scanner.Err()
}
