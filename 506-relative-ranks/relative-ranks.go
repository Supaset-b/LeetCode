func findRelativeRanks(score []int) []string {
    lengthScore := len(score)
    rankScore := make([]string, lengthScore)
    sortScore := make([]int, lengthScore)
    copy(sortScore, score)
    slices.Sort(sortScore)
    for i,s := range score {
        for j,n:= range sortScore {
            if s == n {
                switch j {
                    case lengthScore-1:
                        rankScore[i] = "Gold Medal"
                    case lengthScore-2:
                        rankScore[i] = "Silver Medal"
                    case lengthScore-3:
                        rankScore[i] = "Bronze Medal"
                    default:
                        rankScore[i] = strconv.Itoa(lengthScore-j)
                }
            }
        }
    }
    return rankScore
}