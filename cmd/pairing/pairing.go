package main
import (
    /*
    "log"
    "fmt"
    "encoding/json"
    */
    "github.com/khokhlin/go-tournament-pairing/internal/swiss"
)


func main(){
    swiss.HandleRequests()
}
/*
func main() {

    player1 := swiss.Player{1, 999}
    player2 := swiss.Player{2, 998}
    player3 := swiss.Player{3, 997}
    player4 := swiss.Player{4, 996}
    player5 := swiss.Player{5, 994}
    pairs := swiss.PairPlayers([]swiss.Player{player1, player2, player3, player4, player5})
    for _, p := range pairs {
        fmt.Println(" --- ", p.White.Id, p.Black.Id)
    }

    // Round 1
    game1Result := swiss.GameResult{White: player1,
        Black: player2,
        Result: swiss.WhiteWin}
    game2Result := swiss.GameResult{
        White: player3,
        Black: player4,
        Result: swiss.Draw}
    round1 := swiss.Round{Results: []swiss.GameResult{game1Result, game2Result}}

    // Round 2
    game3Result := swiss.GameResult{
        White: player4,
        Black: player2,
        Result: swiss.WhiteWin}
    game4Result := swiss.GameResult{
        White: player3,
        Black: player1,
        Result: swiss.BlackWin}
    round2 := swiss.Round{Results: []swiss.GameResult{game3Result, game4Result}}

    rounds := []swiss.Round{round1, round2}
    stats := swiss.GetTournamentStats(rounds)
    // fmt.Println(stats)
    marshalled, err := json.Marshal(stats)
    if err != nil {
        log.Fatal("Unable to marshal stats")
    }
    log.Print(string(marshalled))
}
*/
