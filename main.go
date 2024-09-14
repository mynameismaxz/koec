package main

import (
	"github.com/Kong/go-pdk/server"
	"github.com/mynameismaxz/koec/pkg/kong"
)

const (
	Version  = "0.0.1"
	Priority = 1
)

func main() {
	server.StartServer(kong.New, Version, Priority)
}
