func intersect(nums1 []int, nums2 []int) []int {
   lenNum1 := len(nums1)
   lenNum2 := len(nums2)
   m := make(map[int]int)
   index := 0
   if lenNum1 < lenNum2 {
    for _, n1 := range nums1 {
        for i2, n2 := range nums2 {
            if n1 == n2 {
                m[index] = n1
                index++
                nums2 = slices.Delete(nums2, i2, i2+1)
                break;
            }
        }
    }
   } else {
    for _, n2 := range nums2 {
            for i1, n1 := range nums1 {
                if n1 == n2 {
                    m[index] = n1
                    index++
                    nums1 = slices.Delete(nums1, i1, i1+1)
                    break;
                }
            }
        }
   }
   ret := make([]int, len(m))
   for k,v := range m {
    ret[k] = v
   }
   return ret
}