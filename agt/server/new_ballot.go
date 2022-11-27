package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ia04/agt"
	"ia04/comsoc"
	"net/http"
	"time"
)

func (*BallotAgent) decodeNewBallotRequest(r *http.Request) (req agt.NewBallotRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func isValidRule(rule string) bool {
	switch rule {
	case
		"majority",
		"borda",
		"approval":
		return true
	}
	return false
}

func (ba *BallotAgent) doNewBallot(w http.ResponseWriter, r *http.Request) {
	ba.Lock()
	defer ba.Unlock()

	// vérification de la méthode de la requête
	if !ba.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := ba.decodeNewBallotRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// check vote rule
	if !isValidRule(req.Rule) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, "vote rule not implemented")
		return
	}

	// check if the deadline is valid
	if req.Deadline != "" {
		layout := "Mon Jan 02 15:04:05 MST 2006"
		deadline, err := time.Parse(layout, req.Deadline)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
			return
		}
		if deadline.Before(time.Now()) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "deadline is in the past")
			return
		}
	}

	// check if number of alts is valid
	if req.NumberAlts <= 0 {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, "incorrect number of alts")
		return
	}

	// initialise les votes
	voters := make(map[string][]comsoc.Alternative, len(req.VoterIds))
	for _, voterId := range req.VoterIds {
		voters[voterId] = nil
	}

	// initialise les options des votes
	options := make(map[string][]int, len(req.VoterIds))
	for _, voterId := range req.VoterIds {
		options[voterId] = nil
	}

	alts := make([]comsoc.Alternative, req.NumberAlts)

	// initialize alts
	var i comsoc.Alternative
	i = 0
	for range alts {
		alts[i] = i + 1
		i++
	}

	// traitement de la requête
	id := fmt.Sprintf("vote%d", len(ba.ballots)+1)
	ballot := Ballot{
		id:        id,
		rule:      req.Rule,
		alts:      alts,
		deadline:  req.Deadline,
		voters:    voters,
		options:   options,
		voteCount: 0,
	}
	ba.ballots[id] = &ballot

	resp := agt.NewBallotResponse{BallotId: id}
	w.WriteHeader(http.StatusCreated)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
