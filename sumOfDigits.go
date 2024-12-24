
package main
import "fmt"

func sumOfDigits(n int) int {
    sum := 0
    for n != 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    fmt.Printf("Sum of digits of %d is %d
", num, sumOfDigits(num))
}
