package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pgiberti/ia04/agt"
	"gitlab.utc.fr/pgiberti/ia04/comsoc"
	"net/http"
	"time"
)

func (*BallotAgent) decodeVoteRequest(r *http.Request) (req agt.VoteRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (ba *BallotAgent) doVote(w http.ResponseWriter, r *http.Request) {
	// mise à jour du nombre de requêtes
	ba.Lock()
	defer ba.Unlock()

	// vérification de la méthode de la requête
	if !ba.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := ba.decodeVoteRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// check if the vote-id is correct
	if _, exists := ba.ballots[req.VoteId]; exists {
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "incorrect vote-id")
		return
	}

	// check if the deadline has passed
	if ba.ballots[req.VoteId].deadline != "" {
		layout := "Mon Jan 02 15:04:05 MST 2006"
		deadline, _ := time.Parse(layout, ba.ballots[req.VoteId].deadline)
		if deadline.Before(time.Now()) {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "the deadline has passed")
			return
		}
	}

	// check if the agent is in the list
	if val, exists := ba.ballots[req.VoteId].voters[req.AgentId]; exists {
		// check if vote isn't already done
		if val != nil {
			// vote already done
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "this agent has already voted")
			return
		}
	} else {
		// agent not in the list
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "this agent is not in the list")
		return
	}

	// check alts in vote
	err = comsoc.CheckVoteAlternative(req.Prefs, ba.ballots[req.VoteId].alts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// check if prefs isn't empty
	if req.Prefs == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "no preferences given")
		return
	}

	// check if threshold is provided when the method is approval voting
	if ba.ballots[req.VoteId].rule == "approval" {
		if len(req.Options) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "no threshold given")
			return
		}
		if (req.Options[0] < 1) || (req.Options[0] > len(ba.ballots[req.VoteId].alts)) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "incorrect threshold")
			return
		}
	}

	// mets à jour le profil
	fmt.Println("voter", req.AgentId, "voted on ballot", req.VoteId)

	ba.ballots[req.VoteId].voters[req.AgentId] = req.Prefs
	ba.ballots[req.VoteId].options[req.AgentId] = req.Options

	ba.ballots[req.VoteId].voteCount++

	// traitement de la requête
	w.WriteHeader(http.StatusOK)
}
