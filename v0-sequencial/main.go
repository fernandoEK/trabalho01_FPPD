package main

import (
	"fmt"
	"time"
)

func Philosopher(id, iteration int) {
	fmt.Printf("P%d | rodada %d | pensando\n", id+1, iteration)
	time.Sleep(2 * time.Second)

	fmt.Printf("P%d | rodada %d | comendo\n", id+1, iteration)
	time.Sleep(3 * time.Second)
}

func main() {
	philosophers := 5
	rounds := 20

	start := time.Now()

	for i := 1; i <= rounds; i++ {
		for j := 0; j < philosophers; j++ {
			Philosopher(j, i)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\nTempo total da versão sequencial: %s\n", elapsed)
}
