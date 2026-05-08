func findDisappearedNumbers(nums []int) []int {
    l := 0
    r := len(nums)-1
    m := make(map[int]bool)
    for l<=r{
        m[nums[l]] = true
        m[nums[r]] = true
        l++
        r--
    }
    var missNums []int 
    for i:=1;i<=len(nums);i++ {
        if m[i] == false {
            missNums = append(missNums, i)
        }
    }
    return missNums
}