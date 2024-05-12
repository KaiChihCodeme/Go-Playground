package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	cnt := 100
	wg.Add(cnt)
	for _ = range cnt {
		go DoTask(wg)
	}

	wg.Wait()
	// repeatly use the wg
	wg.Add(1)
	go DoFinal(wg)
	wg.Wait()
}

func DoTask(wg *sync.WaitGroup) {
	defer wg.Done()
	return
}

func DoFinal(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("All wg done!")
	return
}
