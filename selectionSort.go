
package main
import "fmt"

func selectionSort(arr []int) {
    for i := 0; i < len(arr); i++ {
        minIdx := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

func main() {
    arr := []int{5, 2, 9, 1, 5, 6}
    fmt.Println("Unsorted array:", arr)
    selectionSort(arr)
    fmt.Println("Sorted array:", arr)
}
