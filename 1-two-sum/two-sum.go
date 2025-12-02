func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i,num := range nums {
        findSecondTarget := target-num
        if index, found := numMap[findSecondTarget]; found {
            return []int{index, i} 
        }
        numMap[num] = i
    }
    return nil
}
