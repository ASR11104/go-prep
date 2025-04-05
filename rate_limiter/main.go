package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

type Limter struct {
	mu          sync.Mutex
	capacity    int
	token       int
	lastUpdated time.Time
	window      int
}

func (l *Limter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(l.lastUpdated)
	if elapsed > time.Duration(l.window)*time.Second {
		l.token = l.capacity
		l.lastUpdated = now
	}
	if l.token > 0 {
		l.token--
		return true
	}
	return false
}

func NewRateLimiter(capacity int, window int) *Limter {
	var limiter *Limter
	once.Do(func() {
		limiter = &Limter{
			capacity:    capacity,
			token:       capacity,
			lastUpdated: time.Now(),
			window:      window,
		}
	})
	return limiter
}

func main() {
	newLimiter := NewRateLimiter(10, 5)
	for i := 0; i < 100; i++ {
		if i%15 == 0 {
			time.Sleep(time.Duration(5) * time.Second)
		}
		if newLimiter.Allow() {
			fmt.Println("Requestion Allowed", i)
		} else {
			fmt.Println("Requestion Denied", i)
		}
	}
}
