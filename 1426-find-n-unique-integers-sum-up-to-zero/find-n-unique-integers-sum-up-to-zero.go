func sumZero(n int) []int {
    ret := make([]int, n)
    i,j := 0,1
    if n%2 != 0{
        ret[0] = 0
        i++
    }
    for i<n {
        ret[i] = j
        ret[i+1] = j*(-1)
        i += 2
        j += 2
    }
    return ret
}