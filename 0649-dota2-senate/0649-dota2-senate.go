func predictPartyVictory(senate string) string {
    partyCount := make(map[rune]int)
    for _, p := range senate {
        partyCount[p]++
    }

    rVote, dVote := 0, 0
    for partyCount['R'] != 0 && partyCount['D'] != 0 {
        if senate[0] == 'R' {
            if dVote == 0 {
                rVote++
                senate = senate[1:] + "R"
            } else {
                dVote--
                partyCount['R']--
                senate = senate[1:]
            }
        } else {
            if rVote == 0 {
                dVote++
                senate = senate[1:] + "D"
            } else {
                rVote--
                partyCount['D']--
                senate = senate[1:]
            }
        }
    }
    if senate[0] == 'R' {
        return "Radiant"
    }
    return "Dire"
}