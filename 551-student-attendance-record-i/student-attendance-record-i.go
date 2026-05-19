func checkRecord(s string) bool {
    cntA:=0
    cntL:=0
    for _,r := range s {
        if r == 'A' {
            cntA++
            cntL=0
            if cntA==2 {
                return false
            }
        } else if r == 'L' {
            cntL++
            if cntL ==3 {
                return false
            }
        } else {
            cntL=0
        }
    }
    return true
}