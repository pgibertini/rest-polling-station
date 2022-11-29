package main

import (
	"fmt"
	"github.com/pgibertini/rest-polling-station/agt/server"
)

func main() {
	serv := server.NewBallotAgent(":8080")
	serv.Start()
	fmt.Scanln()
}
