
package main
import "fmt"

func bubbleSort(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    arr := []int{5, 2, 9, 1, 5, 6}
    fmt.Println("Unsorted array:", arr)
    bubbleSort(arr)
    fmt.Println("Sorted array:", arr)
}
