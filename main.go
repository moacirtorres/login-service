package main

import (
	"github.com/hyperyuri/webapi-with-go/server"
)

func main() {

	s := server.NewServer()

	s.Run()
}
