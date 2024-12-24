
package main
import "fmt"

func printPattern(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    printPattern(num)
}
