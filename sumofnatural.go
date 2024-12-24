
package main
import "fmt"

func sumOfNatural(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    fmt.Printf("Sum of first %d natural numbers is %d
", num, sumOfNatural(num))
}
