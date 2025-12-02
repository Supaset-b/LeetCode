func twoSum(nums []int, target int) []int {
    myMap := make(map[int]int)
    for i,num := range nums {
        findSecondTarget := target-num
        if index, found := myMap[findSecondTarget]; found {
            return []int{index, i} 
        }
        myMap[num] = i
    }
    return nil
}
