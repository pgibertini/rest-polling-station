package main

import (
	"fmt"
	"ia04/agt/server"
)

func main() {
	serv := server.NewBallotAgent(":8080")
	serv.Start()
	fmt.Scanln()
}
