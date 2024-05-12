package main

import (
	"fmt"
	"sync"
	"time"
)

// In a concurrent web server, we assume 100 IPs reach to server at the same time and call server 1000 times repeatly.
// We want to limit the max interval time in 3 min each IP.
// so the final result only 100, not 1000*100
type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan() *Ban {
	return &Ban{
		visitIPs: make(map[string]time.Time),
	}
}

func (b *Ban) visited(ip string) bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	if _, ok := b.visitIPs[ip]; ok {
		return true
	}

	b.visitIPs[ip] = time.Now()
	return false
}

func main() {
	success := 0
	ban := NewBan()
	wg := &sync.WaitGroup{}
	wg.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visited(ip) {
					success += 1
				}
			}()
		}
	}
	wg.Wait()
	fmt.Print(success)
}
