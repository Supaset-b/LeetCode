func reverseBits(n int) int {
    // reverseN := 0
    // index := -1
    // for n != 1 {
    //     // increase index by 1
    //     index++

    //     // % mod for get num
    //     b := n % 2
        
    //     // reverse bit
    //     if b == 0 {
    //         b = 1
    //         // convert bit to nums
    //         reverseN += int(math.Pow( float64(2), float64(index)))
    //     } else {
    //         b = 0
    //     }

    //     // / divide for continue next
    //     n = n/2
    // }

    // return reverseN
    
    // create string fill 0 for 32 bits
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