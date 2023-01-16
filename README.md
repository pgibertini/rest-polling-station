# Online polling station
### Pierre Gibertini, Amaury Michel

Online polling station using Go and REST server.
This project was realized in the context of the IA04 course "Multi-agent systems" at the University of Technology of Compiègne (France).

***

## Installation
#### 1. with `go install`
`go install -v github.com/pgibertini/rest-polling-station@latest` to install 

`$GOPATH/bin/rest-polling-station` (or similar) to launch the server on local port `8080` using the created binary file from your `$GOPATH`

#### 2. with `git`
`git clone https://github.com/pgibertini/rest-polling-station.git` to clone repo

`go run launch_ballotagent.go` to launch the server on local port `8080`


## Work done
- Complete implementation of the [requested API](https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/docs/sujets/activit%C3%A9s/serveur-vote/api.md), with all the necessary verifications
- Implementation of following vote rules : `majority`, `borda` and `approval`


## Implementation choices
- And empty string (`""`) in the field `deadline` for a request `/new_ballot` is accepted, in order to do testing. In this case, you can access the results at any time, and continue voting after accessing the results.
- If the date given for the `deadline` field in a `/new_ballot` query is earlier than `date.Now()`, the ballot is not created.
- Not all registered voters are required to vote. The final result depends on only those agents who voted.
- In case of a tie, a default *tie-break* has been implemented: the candidate with the lowest number is elected.
- The `ranking` given by a `result` query displays the tied candidates in a random order, but necessarily after the candidates with more votes. *On second thought, it would have been better to take into account the tie-break method*.
- The script allowing to launch the voters is very basic and not very functional because we realized most of our tests with the extension `RESTED` of `Mozilla Firefox`.


## Requested web service API

### Command `/new_ballot`

- Request : `POST`
- `JSON` object sent

| property  | type        | example of possible values                                    |
|------------|-------------|-------------------------------------------------------------------|
| `rule`      | `string`       | `"majority"`,`"borda"`, `"approval"`, `"stv"`, `"kemeny"`,... |
| `deadline`  | `string`       | `"Tue Nov 10 23:00:00 UTC 2009"`                               |
| `voter-ids` | `[string,...]` | `["ag_id1", "ag_id2", "ag_id3"]`                                       |
| `#alts`     | `int`          | `12` |   

*Notes:* the deadline represents the end date of the vote. For this, use the standard `Go` library, in particular the `time` package. The `#alts` property represents the number of alternatives, numbered from 1 to `#alts`.

- return code

| return code | meaning |
|-------------|---------------|
| `201`       | ballot created    |
| `400`       | bad request   |
| `501` 	  | not implemented |

- `JSON` object returned (si `201`)

| property  | type | example of possible values                              |
|------------|-------------|-----------------------------------------------------|
| `ballot-id`    | `string` | `"vote12"` |

### Command `/vote`

- Request : `POST`
- `JSON` object sent

| property   | type | example of possible values  |
|------------|-------------|------------------------|
| `agent-id` | `string` | `"ag_id1"` |
| `vote-id`  | `string` | `"vote12"` |
| `prefs`    | `[int,...]` | `[1, 2, 4, 3]` |
| `options`  | `[int,...]` | `[3]` |

*Note: *`options` is optional and allows additional information to be passed (e.g. approval threshold)

- return code

| return code | meaning |
|-------------|---------------|
| `200`       | vote taken into account  |
| `400`       | bad request          |
| `403`       |	vote already taken   |
| `501` 	    | Not Implemented      |
| `503`       | the deadline has passed |

### Commande `/result`

- Request : `POST`
- `JSON` object sent

| property  | type | example of possible values                                 |
|------------|-------------|-----------------------------------------------------|
| `ballot-id`    | `string` | `"vote12"` |


- return code

| return code | meaning   |
|-------------|-----------------|
| `200`       | OK              |
| `425`       | Too early       |
| `404`       |	Not Found       |

- `JSON` object returned (if `200`)

| property   | type | example of possible values  |
|------------|-------------|------------------------|
| `winner`   | `int`       | `4`                    |
| `ranking`  | `[int,...]` | `[2, 1, 4, 3]`         |
