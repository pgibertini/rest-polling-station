package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pgibertini/rest-polling-station/agt"
	"github.com/pgibertini/rest-polling-station/comsoc"
	"net/http"
	"time"
)

func (*BallotAgent) decodeResultRequest(r *http.Request) (req agt.ResultRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (ba *BallotAgent) doResult(w http.ResponseWriter, r *http.Request) {
	ba.Lock()
	defer ba.Unlock()

	// vérification de la méthode de vote
	if !ba.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := ba.decodeResultRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}

	// check if the vote-id is correct
	if _, exists := ba.ballots[req.BallotId]; exists {
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "incorrect ballot-id")
		return
	}

	// check if the deadline has passed
	if ba.ballots[req.BallotId].deadline != "" {
		layout := "Mon Jan 02 15:04:05 MST 2006"
		deadline, _ := time.Parse(layout, ba.ballots[req.BallotId].deadline)
		if deadline.After(time.Now()) {
			w.WriteHeader(http.StatusTooEarly)
			fmt.Fprint(w, "voting is not yet finished")
			return
		}
	}

	// check if voteCount is not 0 (nobody has voted)
	if ba.ballots[req.BallotId].voteCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "no vote found")
		return
	}

	// création du profil à partir des votes effectués (on néglige les personnes n'ayant pas voté)
	profile := make(comsoc.Profile, ba.ballots[req.BallotId].voteCount)
	thresholds := make([]int, ba.ballots[req.BallotId].voteCount)

	i := 0
	for voterId, vote := range ba.ballots[req.BallotId].voters {
		if vote != nil {
			profile[i] = vote

			// save options if vote rule is "approval"
			if ba.ballots[req.BallotId].rule == "approval" {
				thresholds[i] = ba.ballots[req.BallotId].options[voterId][0]
			}
			i++
		}
	}

	var count comsoc.Count
	var winners []comsoc.Alternative
	var ranking []comsoc.Alternative

	switch ba.ballots[req.BallotId].rule {
	case "majority":
		count, err = comsoc.MajoritySWF(profile)
		winners, err = comsoc.MajoritySCF(profile)
	case "borda":
		count, err = comsoc.BordaSWF(profile)
		winners, err = comsoc.BordaSCF(profile)
	case "approval":
		count, err = comsoc.ApprovalSWF(profile, thresholds)
		winners, err = comsoc.ApprovalSCF(profile, thresholds)
	default:
		// n'est pas censé arriver car vérification faite à la création
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, "unknown vote method")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, err.Error())
		return
	}

	winner, err := comsoc.MinTieBreak(winners)

	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, err.Error())
		return
	}

	ranking = comsoc.Ranking(count)

	resp := agt.ResultResponse{Winner: winner, Ranking: ranking}

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
