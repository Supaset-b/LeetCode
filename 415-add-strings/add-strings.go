func addStrings(num1 string, num2 string) string {

    offsetNum1:=len(num1)-1
    offsetNum2:=len(num2)-1
    extender:=0
    sum:=""

    for offsetNum1 >= 0 || offsetNum2 >=0 || extender > 0{
        n1:=0
        n2:=0
        if offsetNum1 >=0 {
            n1 = int(num1[offsetNum1]) - '0'
            offsetNum1--
        }
        
        if offsetNum2 >=0 {
            n2 = int(num2[offsetNum2]) - '0'
            offsetNum2--
        }

        sumTemp:= n1 + n2 + extender
        sum = strconv.Itoa(sumTemp%10) + sum
        extender = sumTemp/10
    }
    return sum
}