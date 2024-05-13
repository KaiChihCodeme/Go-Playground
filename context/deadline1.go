package main

import (
	"fmt"
	"context"
	"time"
)

// Create 3 subtask by go routine, set up a deadline and kill it if time is up

func main() {
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 2)
	go task(ctx)
	time.Sleep(time.Second * 5)
}

func task(ctx context.Context) {
	fmt.Println("tasks start")
	go subtask(ctx)
	go subtask(ctx)
	go subtask(ctx)
}

func subtask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("task timeout")
			return
		default:
			fmt.Println("subtask working")
			time.Sleep(time.Second * 3)
		}	
	}
}