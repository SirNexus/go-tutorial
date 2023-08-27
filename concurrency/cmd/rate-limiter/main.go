package main

import (
	"log"

	"github.com/SirNexus/go-tutorial/concurrency/internal/cmd/ratelimiter"
)

func main() {
	if err := ratelimiter.Cmd().Execute(); err != nil {
		log.Fatalf("could not run rate limiter command")
	}
}
