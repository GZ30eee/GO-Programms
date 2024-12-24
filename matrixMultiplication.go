
package main
import "fmt"

func multiplyMatrices(a, b [][]int) [][]int {
    rowsA := len(a)
    colsA := len(a[0])
    rowsB := len(b)
    colsB := len(b[0])

    if colsA != rowsB {
        fmt.Println("Matrix multiplication not possible")
        return nil
    }

    result := make([][]int, rowsA)
    for i := range result {
        result[i] = make([]int, colsB)
    }

    for i := 0; i < rowsA; i++ {
        for j := 0; j < colsB; j++ {
            for k := 0; k < colsA; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return result
}

func main() {
    matrixA := [][]int{{1, 2}, {3, 4}}
    matrixB := [][]int{{5, 6}, {7, 8}}

    result := multiplyMatrices(matrixA, matrixB)

    fmt.Println("Result of Matrix Multiplication:")
    for _, row := range result {
        fmt.Println(row)
    }
}
