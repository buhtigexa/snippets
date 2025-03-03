package main

import (
	"fmt"
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
	"net/http"
)

func (a *Application) authMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, r.UserAgent())
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (a *Application) rateLimitMe(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		clientKey := realip.FromRequest(r)
		a.mu.Lock()
		if v, ok := a.limiter[clientKey]; !ok {
			a.limiter[clientKey] = rate.NewLimiter(rate.Limit(cfg.rate), cfg.bucket)
		} else {
			if !v.Allow() {
				a.mu.Unlock()
				http.Error(w, http.StatusText(429), 429)
				return
			}
		}
		a.mu.Unlock()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (a *Application) panicker(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := recover(); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Connection", "close")
			return

		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
