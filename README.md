# Online polling station
### Pierre Gibertini, Amaury Michel

Online polling station using Go and REST server.
***

### Installation
#### 1. with `go install`
`go install -v github.com/pgibertini/rest-polling-station@latest` to install 

`$GOPATH/bin/ia04` to launch the server on local port 8080


#### 2. with `git`
`git clone https://github.com/pgibertini/rest-polling-station.git` to clone repo

`go run launch_ballotagent.go` to launch the server on local port 8080


### Work done
- Complete implementation of the [requested API](https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/docs/sujets/activit%C3%A9s/serveur-vote/api.md), with all the necessary verifications
- Implementation of following vote rules : `majority`, `borda` and `approval`


### Implementation choices
- And empty string (`""`) in the field `deadline` for a request `/new_ballot` is accepted, in order to do testing. In this case, you can access the results at any time, and continue voting after accessing the results.
- If the date given for the `deadline` field in a `/new_ballot` query is earlier than `date.Now()`, the ballot is not created.
- Not all registered voters are required to vote. The final result depends on only those agents who voted.
- In case of a tie, a default *tie-break* has been implemented: the candidate with the lowest number is elected.
- The `ranking` given by a `result` query displays the tied candidates in a random order, but necessarily after the candidates with more votes. *On second thought, it would have been better to take into account the tie-break method*.
- The script allowing to launch the voters is very basic and not very functional because we realized most of our tests with the extension `RESTED` of `Mozilla Firefox`.
