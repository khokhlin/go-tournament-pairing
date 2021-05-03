package swiss

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "bytes"
)
func TestTournamentPairingApi(t *testing.T){
    input := `{"Players": [{"Id": 1, "Rating": 22}, {"Id": 2, "Rating": 33}, {"Id": 3, "Rating": 44}, {"Id": 4, "Rating": 55}], "Rounds": []}`
    var jsonStr = []byte(input)
    req, err := http.NewRequest(http.MethodPost, "/tournament-pairing/", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(pairPlayersHandler)
    handler.ServeHTTP(responseRecorder, req)

    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("Handler status is wrong: got %v wont %v", status, http.StatusOK)
    }
    var pairs []Pair
    err = json.NewDecoder(responseRecorder.Body).Decode(&pairs)
    if len(pairs) != 2 {
        t.Errorf("Pairs: wont 2, got %d", len(pairs))
    }
    firstPair := pairs[0]
    if !(firstPair.White.Id == 4 || firstPair.Black.Id == 4) {
        t.Errorf("Player 4 should be in the first pair")
    }
    lastPair := pairs[len(pairs)-1]
    if !(lastPair.White.Id == 1 || lastPair.Black.Id == 1) {
        t.Errorf("Player 1 should be in the last pair")
    }
}
