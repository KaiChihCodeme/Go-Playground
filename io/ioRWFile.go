package main

import (
	"fmt"
	"os"
)

func main() {
	// read
	b, err := os.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)

	err = os.WriteFile("./output.txt", b, 0644) // permission -> 0 no sudo permision, 6 -> owner rw, 4 -> group r, 4 -> others r
	if err != nil {
		panic(err)
	}
}