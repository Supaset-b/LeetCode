func reverseBits(n int) int {
    reverseN := 0
    for i:=31; i>=0; i-- {
        b := n%2
        if b == 1 {
            reverseN += int(math.Pow( float64(2), float64(i)))
        }
        n = n/2
    }
    return reverseN
}