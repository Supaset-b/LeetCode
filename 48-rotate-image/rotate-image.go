func rotate(matrix [][]int) {
    lenMartrix := len(matrix)
    tempMatrix := make([][]int, lenMartrix)
    // copy(tempMatrix, matrix)
    
    for i := range matrix {
        tempMatrix[i] = make([]int, len(matrix[i]))
        copy(tempMatrix[i], matrix[i])
    }
    for i, _ := range tempMatrix {
        position := lenMartrix-1
        for j:=0; j<lenMartrix; j++ {
            matrix[i][j] = tempMatrix[position][i]
            position--
        }
    }
}