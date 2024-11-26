package bingxgo

import (
	"sync"
	"time"
)

type RateLimiter struct {
	requests map[string]*time.Timer
	mu       sync.Mutex
}

func (r *RateLimiter) Wait(endpoint string) {
	r.mu.Lock()
	if timer, exists := r.requests[endpoint]; exists {
		r.mu.Unlock()
		<-timer.C
	} else {
		r.mu.Unlock()
	}
}
