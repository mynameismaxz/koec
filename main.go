package main

import (
	"os"

	"github.com/Kong/go-pdk/server"
	"github.com/mynameismaxz/koec/pkg/kong"
)

const (
	Version  = "0.0.1"
	Priority = 1
)

func main() {
	if err := server.StartServer(kong.New, Version, Priority); err != nil {
		os.Exit(1)
	}
}
