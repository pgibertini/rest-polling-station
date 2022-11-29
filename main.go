package main

import (
	"fmt"
	srv "github.com/pgibertini/rest-polling-station/agt/server"
)

func main() {
	server := srv.NewBallotAgent(":8080")
	server.Start()
	fmt.Scanln()
}
