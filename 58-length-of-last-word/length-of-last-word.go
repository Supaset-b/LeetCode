func lengthOfLastWord(s string) int {
    arrWord := [1000]string{""}
    indexArr := 0
    cntWhitespace := 0
    for _,n := range s {
        if n == ' '{
            indexArr++
            cntWhitespace++
        } else {
            arrWord[indexArr] = arrWord[indexArr]+string(n)
            cntWhitespace=0
        }
    }
    return len(arrWord[indexArr-cntWhitespace])
}