func twoSum(nums []int, target int) []int {
    sum := make([]int, 2)
    numMap := make(map[int]int)
    for i,num := range nums {
        findSecondTarget := target-num
        if index, found := numMap[findSecondTarget]; found {
            return []int{index, i} 
            // sum[0] = index
            // sum[1] = i
        }
        numMap[num] = i
    }
    return sum
}
