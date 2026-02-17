func repeatedNTimes(nums []int) int {
    mapN := make(map[int]bool)
    for _,n := range nums {
        if mapN[n] {
            return n
        }
        mapN[n] = true
    }
    return -1
}