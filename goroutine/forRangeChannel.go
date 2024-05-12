package main

import (
	"fmt"
	"math/rand"
)

func generateNums(n chan int) {
	count := 0
	for {
		if count != 10{
			num := rand.Intn(10)
			n <- num
			count ++
		} else {
			close(n)  // 只要不需要寫入了，最好close掉，避免for loop deadlock
			break
		}
	}
}

func main() {
	// use only 1 cap channel to get the value by blocking once it have value
	intChan := make(chan int, 1)
	go generateNums(intChan)

	for num := range intChan {
		fmt.Println(num)
	}

	// or 
	// for {
	// 	if num, ok := <-intChan; ok {
	// 		fmt.Println(num)
	// 	} else {
	// 		break
	// 	}
	// }
}