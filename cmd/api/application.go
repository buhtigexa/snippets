package main

import (
	"golang.org/x/time/rate"
	"sync"
)

type config struct {
	port   int
	bucket int
	rate   int
}

type Application struct {
	cfg     config
	limiter map[string]*rate.Limiter
	mu      sync.Mutex
}
