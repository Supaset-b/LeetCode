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
