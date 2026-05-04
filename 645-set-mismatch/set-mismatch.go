func findErrorNums(nums []int) []int {

    twice, missing := 0, 0
    m := make(map[int]int)

    // Count occurrences
    for _, n := range nums {
        m[n]++
    }

    // Find duplicate and missing
    for i := 1; i <= len(nums); i++ {
        if m[i] == 2 {
            twice = i
        }
        if m[i] == 0 {
            missing = i
        }
    }

    return []int{twice, missing}
}