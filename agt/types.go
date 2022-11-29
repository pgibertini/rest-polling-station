package agt

import (
	"github.com/pgibertini/rest-polling-station/comsoc"
)

type NewBallotRequest struct {
	Rule       string   `json:"rule"`
	Deadline   string   `json:"deadline"`
	VoterIds   []string `json:"voter-ids"`
	NumberAlts int      `json:"#alts"`
}

type NewBallotResponse struct {
	BallotId string `json:"ballot-id"`
}

type VoteRequest struct {
	AgentId string               `json:"agent-id"`
	VoteId  string               `json:"vote-id"`
	Prefs   []comsoc.Alternative `json:"prefs"`
	Options []int                `json:"options"`
}

type ResultRequest struct {
	BallotId string `json:"ballot-id"`
}

type ResultResponse struct {
	Winner  comsoc.Alternative   `json:"winner"`
	Ranking []comsoc.Alternative `json:"ranking"`
}
