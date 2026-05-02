package findthedegreeofeachvertex

func findDegrees(matrix [][]int) []int {
    var result []int 
    var count int = 0 
    for i := 0 ; i < len(matrix); i ++{
        for j := 0 ; j < len(matrix); j ++{
            if matrix[i][j] == 1{
                count ++
            }
            
            if j == len(matrix) - 1{
                result = append(result, count)
            }
        } 
        count = 0
    }
    return result
}