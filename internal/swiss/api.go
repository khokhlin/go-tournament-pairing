package swiss

import (
    "log"
    "net/http"
    "encoding/json"
)

type Response struct {
    ResponseType string
    Message string
    Data []byte
}


func pairPlayersHandler(w http.ResponseWriter, r *http.Request) {
    var tournament Tournament
    err := json.NewDecoder(r.Body).Decode(&tournament)
    if err != nil {
        http.Error(w, "Can't decode body", http.StatusBadRequest)
        return
    }
    pairs := PairPlayers(tournament)
    log.Print(pairs)
    json.NewEncoder(w).Encode(pairs)
}

func HandleRequests(){
    http.HandleFunc("/tournament-pairing/", pairPlayersHandler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}

