func mirrorDistance(n int) int {
    
    // reverse n
    reverseN := 0
    num := n
    for num!=0 {
       reverseN = reverseN*10 + num%10
       num = num/10
    }

    if n > reverseN {
        return n-reverseN
    } 
    return reverseN-n 
}