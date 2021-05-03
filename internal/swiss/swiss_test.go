package swiss

import (
    "testing"
    "reflect"
)

func TestGetTournamentStats(t *testing.T) {
    player1 := Player{1, 999}
    player2 := Player{2, 998}
    player3 := Player{3, 997}
    player4 := Player{4, 996}

    // Round 1
    round1 := Round{[]GameResult{
            {
                White: player1,
                Black: player2,
                Result: WhiteWin,
            }, {
                White: player3,
                Black: player4,
                Result: Draw,
            },
        },
    }

    // Round 2
    round2 := Round{[]GameResult{
            {
                White: player4,
                Black: player2,
                Result: WhiteWin,
            }, {
                White: player3,
                Black: player1,
                Result: BlackWin,
            },
        },
    }

    rounds := []Round{round1, round2}
    stats := GetTournamentStats(rounds).PlayersStats

    x_opponent_ids := []int{player2.Id, player3.Id}
    opponent_ids := stats[player1.Id].PreviousOpponentIds
    if !reflect.DeepEqual(opponent_ids, x_opponent_ids){
        t.Fatalf("%v opponents are incorect: got %v want %v", player1.Id, opponent_ids, x_opponent_ids)
    }

    x_opponent_ids = []int{player1.Id, player4.Id}
    opponent_ids = stats[player2.Id].PreviousOpponentIds
    if !reflect.DeepEqual(opponent_ids, x_opponent_ids){
        t.Fatalf("%v opponents are incorect: got %v want %v", player1.Id, opponent_ids, x_opponent_ids)
    }
}

func TestPairing(t *testing.T) {
    player1 := Player{1, 999}
    player2 := Player{2, 998}
    player3 := Player{3, 997}
    player4 := Player{4, 996}
    player5 := Player{5, 995}
    players := []Player{player1, player2, player3, player4, player5}

    // Round 1
    round1 := Round{[]GameResult{
            {
                White: player1,
                Black: player2,
                Result: WhiteWin,
            }, {
                White: player3,
                Black: player4,
                Result: Draw,
            },
        },
    }

    // Round 2
    round2 := Round{[]GameResult{
            {
                White: player4,
                Black: player2,
                Result: WhiteWin,
            }, {
                White: player3,
                Black: player1,
                Result: BlackWin,
            },
        },
    }
    rounds := []Round{round1, round2}

    tournament := Tournament{players, rounds}
    pairs := PairPlayers(tournament)
    if len(pairs) != 3 {
        t.Fatalf("Pairs: wont 3, got %d", len(pairs))
    }
}

