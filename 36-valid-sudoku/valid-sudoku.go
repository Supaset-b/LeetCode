func isValidSudoku(board [][]byte) bool {


    box := make([][]byte, 9)

    // chcek column and row
    for i:=0; i<9; i++ {
        mapCheck := make(map[byte]bool)
        for k:=0; k<9; k++ {
            if board[k][i] == byte('.') {
                continue
            } else if mapCheck[board[k][i]] == true {
                 return false
            } else {
                mapCheck[board[k][i]] = true
            }
        }
        clear(mapCheck)
        for j:=0; j<9; j++ {
            if board[i][j] == byte('.') {
                continue
            } else if mapCheck[board[i][j]] == true {
                return false
            } else {
                mapCheck[board[i][j]] = true
                storeBox(j, i, board[i][j], &box)
            }
        }
    }
    return checkNumBox(box)
}

func storeBox(column int, row int, num byte, box *[][]byte) {
    
    boxIndex := (row/3)*3 + (column/3)
    (*box)[boxIndex] = append((*box)[boxIndex], num)

    // if 0 <= row && row <= 2 {
    //     if 0 <= column && column <= 2 {
    //         (*box)[0] = append((*box)[0], num)
    //     } else if 3 <= column && column <= 5 {
    //         (*box)[1] = append((*box)[1], num)
    //     } else if 6 <= column && column <= 9 {
    //         (*box)[2] = append((*box)[2], num)
    //     }
    // } else if 3 <= row && row<= 5 {
    //     if 0 <= column && column <= 2 {
    //         (*box)[3] = append((*box)[3], num)
    //     } else if 3 <= column && column <= 5 {
    //         (*box)[4] = append((*box)[4], num)
    //     } else if 6 <= column && column <= 9 {
    //         (*box)[5] = append((*box)[5], num)
    //     }
    // } else if 6 <= row && row <= 9 {
    //     if 0 <= column && column <= 2 {
    //         (*box)[6] = append((*box)[6], num)
    //     } else if 3 <= column && column <= 5 {
    //         (*box)[7] = append((*box)[7], num)
    //     } else if 6 <= column && column <= 9 {
    //         (*box)[8] = append((*box)[8], num)
    //     }
    // }
}

func checkNumBox(box [][]byte) bool{
   
    for i:=0; i<9; i++{
        mapCheck := make(map[byte]bool)
       for _, val := range box[i]{
            if mapCheck[val] {
                return false
            } else {
                mapCheck[val] = true
            }
        }
    }
    return true
}
