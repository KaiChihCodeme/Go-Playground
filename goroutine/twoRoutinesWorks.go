package main

import (
	"fmt"
	"sync"
)

func main() {
	// print 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
	numFlag, letterFlag := make(chan bool), make(chan bool)
	wg := &sync.WaitGroup{}

	go func() {
		i := 0
		for {
			select {
			case <-numFlag:
				fmt.Print(i)
				i += 1
				fmt.Print(i)
				i += 1
				letterFlag <- true
			}
		}
	}()

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letterFlag:
				if i >= 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(i))
				i += 1
				fmt.Print(string(i))
				i += 1
				numFlag <- true
			}
		}
	}(wg)

	numFlag <- true // Triger the first goroutine
	wg.Wait()       // Wait for the all goroutines to finish
}
