
package main
import "fmt"

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func lcm(a, b int) int {
    return (a * b) / gcd(a, b)
}

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)
    fmt.Printf("LCM of %d and %d is %d
", a, b, lcm(a, b))
}
