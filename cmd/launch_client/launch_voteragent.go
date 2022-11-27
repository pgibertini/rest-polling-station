package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ia04/agt"
	"ia04/agt/client"
	"ia04/comsoc"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	createBallot()

	voter := client.NewVoterAgent("1", "http://localhost:8080", "vote1")
	voter.Vote(randomPrefs(2), nil)

	fmt.Scanln()
}

func createBallot() {
	req := agt.NewBallotRequest{
		Rule:       "majority",
		Deadline:   "",
		VoterIds:   []string{"1", "2"},
		NumberAlts: 2,
	}

	// sérialisation de la requête
	url := "http://localhost:8080/new_ballot"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.St {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.StatusCreated)
		return
	}

	return
}

func randomPrefs(nbCandidats int) []comsoc.Alternative {
	prefs := make([]comsoc.Alternative, nbCandidats)

	for i := 0; i < nbCandidats; i++ {
		prefs[i] = comsoc.Alternative(i + 1)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(prefs), func(i, j int) { prefs[i], prefs[j] = prefs[j], prefs[i] })
	return prefs
}
