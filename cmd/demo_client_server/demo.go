package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pgiberti/ia04/agt"
	"gitlab.utc.fr/pgiberti/ia04/agt/client"
	"gitlab.utc.fr/pgiberti/ia04/agt/server"
	"gitlab.utc.fr/pgiberti/ia04/comsoc"
	"net/http"
	"time"
)

func main() {
	// création du serveur
	serv := server.NewBallotAgent(":8080")
	go serv.Start()
	time.Sleep(time.Second)

	// create a first ballot
	ballotId := createBallot(agt.NewBallotRequest{
		Rule:       "majority",
		Deadline:   "",
		VoterIds:   []string{"a", "b", "c"},
		NumberAlts: 3,
	})

	fmt.Println(ballotId)

	// create voters
	voterA := client.NewVoterAgent("a", "http://localhost:8080", ballotId)
	voterB := client.NewVoterAgent("a", "http://localhost:8080", ballotId)
	voterC := client.NewVoterAgent("a", "http://localhost:8080", ballotId)

	go voterA.Vote([]comsoc.Alternative{1, 2, 3}, nil)
	go voterB.Vote([]comsoc.Alternative{1, 2, 3}, nil)
	go voterC.Vote([]comsoc.Alternative{3, 2, 1}, nil)

	winner := getResult(agt.ResultRequest{BallotId: ballotId})
	fmt.Println(winner)

	fmt.Scanln()
}

func createBallot(req agt.NewBallotRequest) (ballotId string) {
	// sérialisation de la requête
	url := "http://localhost:8080/new_ballot"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	fmt.Println(resp.Header)
	ballotId = resp.Header.Get("ballot-id")
	return
}

func getResult(req agt.ResultRequest) (winner string) {
	// sérialisation de la requête
	url := "http://localhost:8080/new_ballot"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	winner = resp.Header.Get("winner")
	return
}
