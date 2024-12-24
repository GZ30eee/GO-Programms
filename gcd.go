
package main
import "fmt"

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)
    fmt.Printf("GCD of %d and %d is %d
", a, b, gcd(a, b))
}
