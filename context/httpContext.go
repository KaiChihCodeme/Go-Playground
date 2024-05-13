package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Create a simple api server, when in /hello path, setup 10s timeout in this http server
// and write name into context when time is 5s. tehn write this response

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Enter hello")
	defer fmt.Println("end of hello")

	ctx := req.Context()
	// ctx, _ = context.WithTimeout(context.Background(), time.Second * 10)
	ctx, _ = context.WithTimeout(ctx, time.Second * 3)
	ctx = context.WithValue(ctx, "name", "kai")

	res := make(chan string, 1)
	go task(ctx, res) 

	select {
	case <-ctx.Done():
		fmt.Println("Server get err: ", ctx.Err())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	case m := <-res:
		w.Write([]byte(m))
	}

}

func task(ctx context.Context, res chan string) {
	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
	case <-time.After(5 * time.Second):
		res <- fmt.Sprintf("name is : %s", ctx.Value("name"))
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}