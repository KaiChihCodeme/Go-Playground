package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		for _ = range 5 {
			out <- rand.Intn(5)
		}
		close(out) // close channel, only read cannot write
	}()

	go func() {
		defer wg.Done()

		for i := range out { // channel can be read by for loop
			fmt.Println(i)
		}
	}()
	wg.Wait()
}
