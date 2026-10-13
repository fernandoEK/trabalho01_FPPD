package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	philosophers = 5
	cycles       = 20
	thinkTime    = 2 * time.Second
	eatTime      = 3 * time.Second
)

var (
	forks [philosophers]sync.Mutex
	table sync.Mutex
	wg    sync.WaitGroup
)

func philosopher(id int) {
	defer wg.Done()

	left := id
	right := (id + 1) % philosophers

	for round := 1; round <= cycles; round++ {
		fmt.Printf("P%d | round %d | thinking\n", id+1, round)
		time.Sleep(thinkTime)

		table.Lock()
		forks[left].Lock()
		forks[right].Lock()

		fmt.Printf("P%d | round %d | eating\n", id+1, round)
		time.Sleep(eatTime)

		forks[right].Unlock()
		forks[left].Unlock()
		table.Unlock()
	}
}

func main() {
	start := time.Now()

	wg.Add(philosophers)
	for id := 0; id < philosophers; id++ {
		go philosopher(id)
	}
	wg.Wait()

	fmt.Printf("\nV2 total time: %s\n", time.Since(start))
}
