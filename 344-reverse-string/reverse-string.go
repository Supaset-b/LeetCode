func reverseString(s []byte)  {
    length := len(s)
    reverseS := make([]byte, length)
    for index, value := range s {
        reverseS[length-index-1] = value
    }
    copy(s, reverseS)
}