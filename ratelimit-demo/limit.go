package main

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// Create a map to hold the rate limiters for each visitor and a mutex
var visitors = make(map[string]*rate.Limiter)
var mtx sync.Mutex

// Create a rate limiter and add it to the visitors map, using the
// IP address as the key
func addVisitor(ip string) *rate.Limiter {
	// func NewLimiter(r Limit, b int) *Limiter
	// permits you to consume an average of r tokens per second, with a maximum of b tokens in any single 'burst'
	limiter := rate.NewLimiter(2, 5)
	mtx.Lock()
	visitors[ip] = limiter
	mtx.Unlock()
	return limiter
}

// Retrive and return the rate limiter for current user
func getVisitor(ip string) *rate.Limiter {
	mtx.Lock()
	limiter, exists := visitors[ip]
	mtx.Unlock()
	if !exists {
		return addVisitor(ip)
	}

	return limiter
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		limiter := getVisitor(r.RemoteAddr)
		// log
		//fmt.Println("Request from : ", r.RemoteAddr)

		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
