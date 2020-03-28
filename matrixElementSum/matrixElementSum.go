func matrixElementsSum(matrix [][]int) int {
    huntedRoom := 0
    for i:=0;i<len(matrix[0]);i++ {
	//fmt.Println("Row: ",len(matrix), matrix[0][i], i)
        if matrix[0][i] == 0 {
            continue
        }
        for j:=0;j<len(matrix);j++ {
            //fmt.Println("col: ", matrix[j][i], j,i)
	    if matrix[j][i] == 0{
                break
            }
            huntedRoom += matrix[j][i]
        }
    }
    return huntedRoom
}

