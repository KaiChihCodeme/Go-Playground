package main

import (
	"os"
	"fmt"
)

func main() {
	// readFile, err := os.ReadFile("/test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
	}	
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}

	// read
	b1 := make([]byte, 10)
	n1, err := f.Read(b1)
	if err != nil {
		fmt.Println(err)
	}	
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// write
	buf := make([]byte, 1024)
	

}