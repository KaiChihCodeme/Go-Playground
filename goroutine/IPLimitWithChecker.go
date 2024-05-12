package main

import (
	"context"
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

func NewBan(ctx context.Context) *Ban {
	b := &Ban{
		visitIPs: make(map[string]time.Time),
	}

	// create a timer routine to check expire ip
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				b.lock.Lock()
				for k, v := range b.visitIPs {
					if time.Now().Sub(v) >= time.Minute*3 { // over 3 min, delete it
						delete(b.visitIPs, k)
					}
				}
				b.lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()

	return b
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ban := NewBan(ctx)
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
