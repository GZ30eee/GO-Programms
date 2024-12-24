
package main
import "fmt"

func countDigits(n int) int {
    count := 0
    for n != 0 {
        count++
        n /= 10
    }
    return count
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    fmt.Printf("The number %d has %d digits.
", num, countDigits(num))
}
