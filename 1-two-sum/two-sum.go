func twoSum(nums []int, target int) []int {
    sum := make([]int, 2)
    numMap := make(map[int]int)

    

    for i,num := range nums {
        findSecondTarget := target-num

        // for index, s := range nums[i:] {
		//     numMap[s] = index
	    // }

        if index, found := numMap[findSecondTarget]; found {
            sum[0] = index
            sum[1] = i
        }
        numMap[num] = i
    }
    return sum
}
