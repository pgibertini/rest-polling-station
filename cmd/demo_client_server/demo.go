package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pgibertini/rest-polling-station/agt"
	"github.com/pgibertini/rest-polling-station/agt/client"
	"github.com/pgibertini/rest-polling-station/agt/server"
	"github.com/pgibertini/rest-polling-station/comsoc"
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

	fmt.Printf("majority vote : %s\n", ballotId)

	// create voters
	voterA := client.NewVoterAgent("a", "http://localhost:8080", ballotId)
	voterB := client.NewVoterAgent("b", "http://localhost:8080", ballotId)
	voterC := client.NewVoterAgent("c", "http://localhost:8080", ballotId)

	go voterA.Vote([]comsoc.Alternative{1, 2, 3}, nil)
	go voterB.Vote([]comsoc.Alternative{1, 2, 3}, nil)
	go voterC.Vote([]comsoc.Alternative{3, 2, 1}, nil)

	time.Sleep(time.Second)

	winner := getResult(agt.ResultRequest{BallotId: ballotId})
	fmt.Printf("[%s] winner : %d\n", ballotId, winner)

	time.Sleep(time.Second)

	// create a second ballot
	ballotId = createBallot(agt.NewBallotRequest{
		Rule:       "borda",
		Deadline:   "",
		VoterIds:   []string{"a", "b", "c"},
		NumberAlts: 3,
	})

	fmt.Printf("\nborda vote : %s\n", ballotId)

	// create voters
	voterA = client.NewVoterAgent("a", "http://localhost:8080", ballotId)
	voterB = client.NewVoterAgent("b", "http://localhost:8080", ballotId)
	voterC = client.NewVoterAgent("c", "http://localhost:8080", ballotId)

	go voterA.Vote([]comsoc.Alternative{2, 3, 1}, nil)
	go voterB.Vote([]comsoc.Alternative{1, 2, 3}, nil)
	go voterC.Vote([]comsoc.Alternative{2, 1, 3}, nil)

	time.Sleep(time.Second)

	winner = getResult(agt.ResultRequest{BallotId: ballotId})
	fmt.Printf("[%s] winner : %d\n", ballotId, winner)

	time.Sleep(time.Second)

	// create a second ballot
	ballotId = createBallot(agt.NewBallotRequest{
		Rule:       "approval",
		Deadline:   "",
		VoterIds:   []string{"a", "b", "c"},
		NumberAlts: 3,
	})

	fmt.Printf("\napproval vote : %s\n", ballotId)

	// create voters
	voterA = client.NewVoterAgent("a", "http://localhost:8080", ballotId)
	voterB = client.NewVoterAgent("b", "http://localhost:8080", ballotId)
	voterC = client.NewVoterAgent("c", "http://localhost:8080", ballotId)

	go voterA.Vote([]comsoc.Alternative{2, 3, 1}, []int{1})
	go voterB.Vote([]comsoc.Alternative{1, 2, 3}, []int{2})
	go voterC.Vote([]comsoc.Alternative{2, 1, 3}, []int{1})

	time.Sleep(time.Second)

	winner = getResult(agt.ResultRequest{BallotId: ballotId})
	fmt.Printf("[%s] winner : %d\n", ballotId, winner)

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
	if resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		fmt.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	var newBallotResp agt.NewBallotResponse
	json.Unmarshal(buf.Bytes(), &newBallotResp)

	ballotId = newBallotResp.BallotId
	return
}

func getResult(req agt.ResultRequest) (winner comsoc.Alternative) {
	// sérialisation de la requête
	url := "http://localhost:8080/result"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		fmt.Println(err)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	var resultResp agt.ResultResponse
	json.Unmarshal(buf.Bytes(), &resultResp)

	winner = resultResp.Winner
	return
}
