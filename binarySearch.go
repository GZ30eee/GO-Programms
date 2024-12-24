
package main
import "fmt"

func binarySearch(arr []int, x int) int {
    low, high := 0, len(arr)-1
    for low <= high {
        mid := low + (high-low)/2
        if arr[mid] == x {
            return mid
        } else if arr[mid] < x {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Print("Enter a number to search: ")
    var num int
    fmt.Scan(&num)
    result := binarySearch(arr, num)
    if result != -1 {
        fmt.Printf("Element found at index %d
", result)
    } else {
        fmt.Println("Element not found")
    }
}
