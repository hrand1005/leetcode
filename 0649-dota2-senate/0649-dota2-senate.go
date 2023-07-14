func predictPartyVictory(senate string) string {
    rqueue := make([]int, 0, len(senate))
    dqueue := make([]int, 0, len(senate))
    for i, p := range senate {
        if p == 'R' {
            rqueue = append(rqueue, i)
        } else {
            dqueue = append(dqueue, i)
        }
    }
    
    for len(rqueue) != 0 && len(dqueue) != 0 {
        r, d := rqueue[0], dqueue[0]
        rqueue, dqueue = rqueue[1:], dqueue[1:]
        
        if r < d {
            rqueue = append(rqueue, r+len(senate))
        } else {
            dqueue = append(dqueue, d+len(senate))
        }
    }
    
    if len(dqueue) == 0 {
        return "Radiant"
    }
    return "Dire"
}
/*
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
*/