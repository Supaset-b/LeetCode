func merge(nums1 []int, m int, nums2 []int, n int)  {

    mergeNums := make([]int, m+n)
    i,i1,i2 := 0,0,0

    for i1 < m && i2 < n{
        if nums1[i1] < nums2[i2] {
            mergeNums[i] = nums1[i1]
            i1++
        } else if nums1[i1] >= nums2[i2] {
            mergeNums[i] = nums2[i2]
            i2++
        }
        i++
    }
    // remain element
    for i1 < m {
        mergeNums[i] = nums1[i1]
        i++
        i1++
    }
    for i2 < n {
        mergeNums[i] = nums2[i2]
        i++
        i2++
    }
    copy(nums1,mergeNums)
}