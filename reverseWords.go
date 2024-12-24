
package main
import "fmt"
import "strings"

func reverseWords(s string) string {
    words := strings.Fields(s)
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ")
}

func main() {
    var str string
    fmt.Print("Enter a string: ")
    fmt.Scan(&str)
    fmt.Println("Reversed words:", reverseWords(str))
}
