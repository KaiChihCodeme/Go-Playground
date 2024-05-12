package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var supervisor1 sync.WaitGroup
var supervisor2 sync.WaitGroup
var boss sync.WaitGroup

func doWork1() {
	defer supervisor1.Done()
	workDuration := rand.Intn(10)
	time.Sleep(time.Duration(workDuration) * time.Second)
	fmt.Println("Work 1 Done!")
}

func doWork2() {
	defer supervisor2.Done()
	workDuration := rand.Intn(10)
	time.Sleep(time.Duration(workDuration) * time.Second)
	fmt.Println("Work 2 Done!")
}

func wait1() {
	defer boss.Done()
	supervisor1.Wait()
	fmt.Println("All work 1 DONE!")
}

func wait2() {
	defer boss.Done()
	supervisor2.Wait()
	fmt.Println("All work 2 DONE!")
}

func waitForAll() {
	defer boss.Done()
	wait1()
	wait2()
	fmt.Println("All WORK DONE!")
}

func main() {

}
