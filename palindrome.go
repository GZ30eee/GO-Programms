
package main
import "fmt"

func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-i-1] {
            return false
        }
    }
    return true
}

func main() {
    var str string
    fmt.Print("Enter a string: ")
    fmt.Scan(&str)
    if isPalindrome(str) {
        fmt.Println("The string is a palindrome.")
    } else {
        fmt.Println("The string is not a palindrome.")
    }
}
