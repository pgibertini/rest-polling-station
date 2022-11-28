package main

import (
	"fmt"
	"gitlab.utc.fr/pgiberti/ia04/agt/server"
)

func main() {
	serv := server.NewBallotAgent(":8080")
	serv.Start()
	fmt.Scanln()
}
