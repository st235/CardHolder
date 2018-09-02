package main

import (
	"card/server"
	"card/core"
)

func main() {
	core.Parse("response.json")

	server.Register()
}
