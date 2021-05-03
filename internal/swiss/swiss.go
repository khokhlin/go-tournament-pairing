package swiss

import (
    "errors"
    "sort"
    "math/rand"
    "time"
)

type result int
const (
    WhiteWin result = iota
    BlackWin
    Draw
)

type Color = int
const (
    White Color = iota + 1
    Black
)

type SinglePlayerStats struct {
    PreviousOpponentIds []int
    WhiteGameCount int
    BlackGameCount int
    TotalPoints int
}

func RandBool() bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(2) == 1
}

func GetRandomColor() Color {
    if RandBool() {
        return White
    } else {
        return Black
    }
}

func (sps SinglePlayerStats) PreferredColor() Color {
    if sps.WhiteGameCount > sps.BlackGameCount {
        return Black
    } else if sps.WhiteGameCount < sps.BlackGameCount {
        return White
    } else {
        return GetRandomColor()
    }
}

type TournamentStats struct {
    PlayersStats map[int]SinglePlayerStats
}

type Player struct {
    Id int
    Rating int
}

type Pair struct {
    White Player
    Black Player
}

type GameResult struct {
    White Player
    Black Player
    Result result
}

type Round struct {
    Results []GameResult
}

type Tournament struct {
    Players []Player
    Rounds []Round
}

func getNextPlayer(players []Player) (Player, []Player, error) {
    if len(players) == 0 {
        return Player{}, players, errors.New("No players left")
    } else {
        idx := 0
        player := players[idx]
        players := append(players[:idx], players[idx+1:]...)
        return player, players, nil
    }
}

func NumInList(num int, list []int) bool {
    for _, idx := range list {
        if idx == num {
            return true
        }
    }
    return false
}

func getNextOpponent(player Player, playerStats SinglePlayerStats, players []Player) (Player, []Player, error) {
    // Find first player haven't played yet.
    for idx, opponentCandidate := range players {
        if !NumInList(opponentCandidate.Id, playerStats.PreviousOpponentIds) {
            players := append(players[:idx], players[idx+1:]...)
            return opponentCandidate, players, nil
        }

    }
    return Player{}, players, errors.New("No opponents left")
}

func sortPlayers(players []Player, stats TournamentStats) []Player {
    sort.Slice(players, func(i, j int) bool {
        left := players[i]
        right := players[j]
        leftTotalPoints := stats.PlayersStats[left.Id].TotalPoints
        rightTotalPoints := stats.PlayersStats[left.Id].TotalPoints
        if leftTotalPoints == rightTotalPoints {
            return left.Rating > right.Rating
        } else {
            return leftTotalPoints > rightTotalPoints
        }
    })
    return players
}

func makePair(playerStats SinglePlayerStats, player Player, opponent Player) Pair {
    var pair Pair
    if playerStats.PreferredColor() == White {
        pair = Pair{player, opponent}
    } else {
        pair = Pair{opponent, player}
    }
    return pair
}

func PairPlayers(tournament Tournament) []Pair {
    var pairs []Pair
    var err error
    var player Player
    var opponent Player
    players := tournament.Players
    stats := GetTournamentStats(tournament.Rounds)
    unpairedPlayers := sortPlayers(players[:], stats)
    lastIdx := len(unpairedPlayers) - 1
    halfIdx := (lastIdx + 1) / 2
    topHalf := unpairedPlayers[:halfIdx]
    bottomHalf := unpairedPlayers[halfIdx:]
    for {
        player, topHalf, err = getNextPlayer(topHalf)
        if err != nil {
            break
        }
        playerStats := stats.PlayersStats[player.Id]
        opponent, bottomHalf, err = getNextOpponent(player, playerStats, bottomHalf)
        if err != nil {
            // If there are no opponents in the bottom half left, search in the top half.
            opponent, topHalf, err = getNextOpponent(player, playerStats, topHalf)
            if err != nil {
                break
            }
        }
        pairs = append(pairs, makePair(playerStats, player, opponent))
    }

    if len(bottomHalf) > 0 {
        player = bottomHalf[0]
        pairs = append(pairs, Pair{player, Player{}})
    }
    return pairs
}

func GetTournamentStats(rounds []Round) TournamentStats {
    stats := TournamentStats{make(map[int]SinglePlayerStats)}
    for _, round := range rounds {
        for _, gameResult := range round.Results {
            white := gameResult.White
            black := gameResult.Black

            whiteStats, ok := stats.PlayersStats[white.Id]
            if ! ok {
                whiteStats = SinglePlayerStats{}
                whiteStats.WhiteGameCount++
            }

            blackStats, ok := stats.PlayersStats[black.Id]
            if ! ok {
                blackStats = SinglePlayerStats{}
                blackStats.BlackGameCount++
            }

            whiteStats.PreviousOpponentIds = append(whiteStats.PreviousOpponentIds, black.Id)
            blackStats.PreviousOpponentIds = append(blackStats.PreviousOpponentIds, white.Id)
            if gameResult.Result == WhiteWin {
                whiteStats.TotalPoints++
                blackStats.TotalPoints--

            } else if gameResult.Result == BlackWin {
                blackStats.TotalPoints++
                whiteStats.TotalPoints--
            } else {
                // Draw: do nothing
            }
            stats.PlayersStats[white.Id] = whiteStats
            stats.PlayersStats[black.Id] = blackStats
        }
    }
    return stats
}
