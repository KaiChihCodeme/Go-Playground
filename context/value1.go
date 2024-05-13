package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "key", "val")
	res := make(chan string, 1)
	go task(ctx, res)
	fmt.Println(<-res)
	fmt.Println("Get Value Success!")
}

func task(ctx context.Context, res chan string) {
	fmt.Println("get ctx value: ", ctx.Value("key"))
	time.Sleep(time.Second * 3)
	res <- ctx.Value("key").(string)
}