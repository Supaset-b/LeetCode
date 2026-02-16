func hammingWeight(n int) int {
   cntBit := 1
   for n != 1 {
        bit := n%2
        if bit == 1{
            cntBit ++
        }
        n = n/2
   }
   return cntBit
}