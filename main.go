package main

import (
	"fmt"
	srv "gitlab.utc.fr/pgiberti/ia04/agt/server"
)

func main() {
	server := srv.NewBallotAgent(":8080")
	server.Start()
	fmt.Scanln()
}
