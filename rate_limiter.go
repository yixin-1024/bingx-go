package bingxgo

import (
	"sync"
	"time"
)

type RateLimiter struct {
	requests map[string]*time.Timer
	mu       sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]*time.Timer),
	}
}

func (r *RateLimiter) Add(endpoint string, duration time.Duration) {
	r.mu.Lock()
	r.requests[endpoint] = time.AfterFunc(duration, func() {
		delete(r.requests, endpoint)
	})
	r.mu.Unlock()
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
