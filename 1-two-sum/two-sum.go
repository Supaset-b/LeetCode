func twoSum(nums []int, target int) []int {
    myMap := make(map[int]int)
    for i,num := range nums {
        find2nd := target-num
        index, found := myMap[find2nd]; 
        if found {
            return []int{index, i} 
        }
        myMap[num] = i
    }
    return nil
}
