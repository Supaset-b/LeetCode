func findRelativeRanks(score []int) []string {
    rankScore := make([]string, len(score))
    sortScore := make([]int, len(score))
    copy(sortScore, score)
    slices.Sort(sortScore)
    for i,s := range score {
        for j,n:= range sortScore {
            if s == n {
                switch j {
                    case len(sortScore)-1:
                        rankScore[i] = "Gold Medal"
                    case len(sortScore)-2:
                        rankScore[i] = "Silver Medal"
                    case len(sortScore)-3:
                        rankScore[i] = "Bronze Medal"
                    default:
                        rankScore[i] = strconv.Itoa(len(sortScore)-j)
                }
            }
        }
    }
    return rankScore
}