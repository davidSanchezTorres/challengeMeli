package main

import (
	"isMutant/handlers"
)

func main() {
	handlers.ServerListener("3000")
}

// go test -coverprofile="coverage.out" ./...
// go tool cover --func=coverage.out
// go tool cover --html=coverage.out
