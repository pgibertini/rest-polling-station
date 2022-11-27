package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pgiberti/ia04/agt"
	"gitlab.utc.fr/pgiberti/ia04/comsoc"
	"net/http"
)

type VoterAgent struct {
	id       string
	url      string
	ballotID string
}

func NewVoterAgent(id string, url string, ballotID string) *VoterAgent {
	return &VoterAgent{id, url, ballotID}
}

func (va *VoterAgent) Vote(prefs []comsoc.Alternative, options []int) {
	req := agt.VoteRequest{
		AgentId: va.id,
		VoteId:  va.ballotID,
		Prefs:   prefs,
		Options: options,
	}

	// sérialisation de la requête
	url := va.url + "/vote"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		fmt.Println(err)
		return
	}

	fmt.Printf("%s has voted\n", va.id)
	return
}
