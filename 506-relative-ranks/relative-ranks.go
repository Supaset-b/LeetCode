func findRelativeRanks(score []int) []string {
    lengthScore := len(score)
    rankScore := make([]string, lengthScore)
    sortScore := make([]int, lengthScore)
    copy(sortScore, score)
    slices.Sort(sortScore)
    mappingSort := make(map[int]int)
    for j,n:= range sortScore {
        mappingSort[n] = j
    }

    // loop for checking ranks
    for i,s := range score { 
        rank := lengthScore - mappingSort[s]
        switch mappingSort[s] {
            case lengthScore-1:
                rankScore[i] = "Gold Medal"
            case lengthScore-2:
                rankScore[i] = "Silver Medal"
            case lengthScore-3:
                rankScore[i] = "Bronze Medal"
            default:
                rankScore[i] = strconv.Itoa(rank)
        }
    }
    return rankScore
}