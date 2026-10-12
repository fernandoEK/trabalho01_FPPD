package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numFilosofos = 5
	numCiclos    = 20
	tempoPensar  = 2 * time.Second
	tempoComer   = 3 * time.Second
)

var garfos [numFilosofos]sync.Mutex
var grupo sync.WaitGroup

func filosofo(id int) {

	esquerda := id
	direita := (id + 1) % numFilosofos

	for rodada := 1; rodada <= numCiclos; rodada++ {

		fmt.Printf("P%d | rodada %d | pensando\n", id+1, rodada)
		time.Sleep(tempoPensar)

		garfos[esquerda].Lock()
		fmt.Printf("P%d | pegou garfo %d\n", id+1, esquerda+1)

		fmt.Printf("P%d | esperando garfo %d\n", id+1, direita+1)
		garfos[direita].Lock()

		fmt.Printf("P%d | rodada %d | comendo\n", id+1, rodada)
		time.Sleep(tempoComer)

		garfos[direita].Unlock()
		garfos[esquerda].Unlock()
	}

	grupo.Done()
}

func main() {

	inicio := time.Now()

	grupo.Add(numFilosofos)

	for id := 0; id < numFilosofos; id++ {
		go filosofo(id)
	}

	grupo.Wait()

	fmt.Printf("\nTempo total da V1: %s\n", time.Since(inicio))
}
