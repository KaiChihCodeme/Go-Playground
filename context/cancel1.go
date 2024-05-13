package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, done chan struct{}, i int) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("Cancel by ticker")
			done <- struct{}{}
		default:
			fmt.Println(i)
			i--
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{}, 1)
	cntDown := 3
	go worker(ctx, done, cntDown)
	time.Sleep(time.Duration(cntDown) * time.Second) // 在這裡停cntDown秒之後call cancel去取消worker
	cancel()
	<-done
}
