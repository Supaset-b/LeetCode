func reversePrefix(word string, ch byte) string {

    index := -1
    for i:=0; i<len(word); i++ {
        if ch==word[i] {
            index = i
            break
        }
    }

    // doesn't has ch in the world
    if index == -1 {
        return word
    }

    // prefix
    prefix := []byte(word[:index+1])
    
    // reverse prefix
    left := 0
    right := len(prefix)-1

    for left < right {
        temp := prefix[right]
        prefix[right] = prefix[left]
        prefix[left] = temp
        right--
        left++
    }

    return string(prefix) + string(word[index+1:])
}