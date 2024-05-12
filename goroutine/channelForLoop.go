package main

import (
	"fmt"
)

func main() {
	num := 100
	testChan := make(chan int, 100)
	for i:=0; i<100; i++ {
		testChan <- num
	}

	close(testChan)  // 因此必須在讀取之前close, 才能用for loop安全讀取

	for i := range testChan{  // 如果不close(testChan), 將導致for loop不斷嘗試取出testChan進而deadlock
		fmt.Println(i)
	}
}