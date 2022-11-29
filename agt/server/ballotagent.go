package server

import (
	"fmt"
	"github.com/pgibertini/rest-polling-station/comsoc"
	"log"
	"net/http"
	"sync"
	"time"
)

type BallotAgent struct {
	sync.Mutex
	id      string
	addr    string
	ballots map[string]*Ballot
}

type Ballot struct {
	id        string
	rule      string
	deadline  string
	voters    map[string][]comsoc.Alternative
	options   map[string][]int
	alts      []comsoc.Alternative
	voteCount int
}

// NewBallotAgent create a new vote service
func NewBallotAgent(addr string) *BallotAgent {
	return &BallotAgent{id: addr, addr: addr, ballots: make(map[string]*Ballot)}
}

// checkMethod test a method
func (ba *BallotAgent) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

// Start starts the REST server
func (ba *BallotAgent) Start() {
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/new_ballot", ba.doNewBallot)
	mux.HandleFunc("/vote", ba.doVote)
	mux.HandleFunc("/result", ba.doResult)

	// création du serveur http
	s := &http.Server{
		Addr:           ba.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	log.Println("Listening on", ba.addr)
	go log.Fatal(s.ListenAndServe())
}
