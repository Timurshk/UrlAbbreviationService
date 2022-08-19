package main

import (
	"github.com/Timurshk/internal/server"
)

const Host = "localhost"
const Port = "8080"

func main() {
	serv := server.New(Host, Port)
	print(serv)
	serv.Start()
}
