package main

import (
	"flag"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"os"
	"strconv"
)

var cfg *config

func init() {
	d, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	cfg = &config{
		port: d,
	}

	cfg.rate, err = strconv.Atoi(os.Getenv("R"))
	if err != nil {
		panic(err)
	}
	cfg.bucket, err = strconv.Atoi(os.Getenv("B"))
	if err != nil {
		panic(err)
	}

}

func main() {
	flag.IntVar(&cfg.port, "port", cfg.port, "-port=<port number> to listen on")
	flag.Parse()
	app := &Application{cfg: *cfg, limiter: make(map[string]*rate.Limiter)}
	mux := http.NewServeMux()

	mux.Handle("/", app.rateLimitMe(app.authMiddleware(http.HandlerFunc(app.createSnippetController))))
	mux.Handle("/snippet/{id}", app.rateLimitMe(app.authMiddleware(http.HandlerFunc(app.getById))))
	log.Printf("Listening on port: %d", cfg.port)
	if err := http.ListenAndServe(fmt.Sprint(":", cfg.port), mux); err != nil {
		panic(err)
	}

}
