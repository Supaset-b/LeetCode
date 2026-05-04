func judgeCircle(moves string) bool {
    x:=0
    y:=0
    m := make(map[rune]int)
    m['U'] = 1
    m['D'] = -1
    m['R'] = 2
    m['L'] = -2
    
    for _,c := range moves {
        if c=='U' || c=='D'{
            y += m[c]
        } else {
            x += m[c]
        }
    }
    return 0==y && 0==x
}