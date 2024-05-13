package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// guess the number, only give you n seconds to guess 0~100. If time is up, return.

func main() {
	n := 1
	ctx, cancel := context.WithCancel(context.Background())
	res := make(chan string, 1)
	luckyNum := 9
	go guess(ctx, res, luckyNum)
	time.Sleep(time.Duration(n) * time.Second)
	cancel()
	fmt.Println(<-res)
}

func guess(ctx context.Context, res chan string, luckyNum int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time is up, guess failed")
			res <- "failed"
			return
		default:
			guessNum := rand.Intn(1000)
			if guessNum == luckyNum {
				fmt.Printf("Guess Success: %d\n", guessNum)
				res <- "success!"
				return
			} else {
				fmt.Printf("NOT THIS: %d!\n", guessNum)
			}
		}
	}

}
