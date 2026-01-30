package service

import "time"

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(rps int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, rps),
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			for len(rl.tokens) < cap(rl.tokens) {
				rl.tokens <- struct{}{}
			}
		}
	}()

	return rl
}

func (r *RateLimiter) Allow() bool {
	select {
	case <-r.tokens:
		return true
	default:
		return false
	}
}
