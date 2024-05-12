package main

import (
	"fmt"
	"time"
)

// Write a program to ensure that do not stop or exit program by panic
// call proc every second, and ensure you won't exit by panic
func main() {
	go func() {
		t := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-t.C:
				// 如果這邊不再創一個goroutine, defer是在goroutine結束時執行，因此會繼續call proc()引發更多panic
				// 但這些panic並沒有被recover, 因此就是無窮迴圈但不做defer的事
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()

	// 此行可以讓main()維持阻塞，不會有程式結束的一天（除非手動關閉）
	select {}
}

func proc() {
	panic("ok")
}
