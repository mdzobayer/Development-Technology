package main

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Create a map to hold the rate limiters for each visitor and a mutex
var visitors = make(map[string]*visitor)
var mtx sync.Mutex

// Run a background goroutine to remove old entities from the visitors map
func init() {
	go cleanupVisitors()
}

// Create a rate limiter and add it to the visitors map, using the
// IP address as the key
func addVisitor(ip string) *rate.Limiter {
	// func NewLimiter(r Limit, b int) *Limiter
	// permits you to consume an average of r tokens per second, with a maximum of b tokens in any single 'burst'
	limiter := rate.NewLimiter(2, 5)
	mtx.Lock()
	visitors[ip] = &visitor{limiter, time.Now()}
	mtx.Unlock()
	return limiter
}

// Retrive and return the rate limiter for current user
func getVisitor(ip string) *rate.Limiter {
	mtx.Lock()
	v, exists := visitors[ip]

	if !exists {
		mtx.Unlock()
		return addVisitor(ip)
	}

	// Update the last seen time for the visitor
	v.lastSeen = time.Now()
	mtx.Unlock()

	return v.limiter
}

// Every minute check the map for visitors that haven't been seen
// more than 3 minute and delete the entries
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		mtx.Lock()
		for ip, v := range visitors {
			if time.Now().Sub(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mtx.Unlock()
	}
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
