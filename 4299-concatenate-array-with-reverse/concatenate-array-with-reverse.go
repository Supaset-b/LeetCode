func concatWithReverse(nums []int) []int {
    reverseNum := make([]int, len(nums)*2)
    r := len(reverseNum)-1
    for i,n := range nums {
        reverseNum[i] = n
        reverseNum[r] = n
        r--
    }
    return reverseNum
}