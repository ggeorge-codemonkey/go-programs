package main

import (
    "fmt"
    "strings"
)

func isPalindrome(s string) bool {
    // Convert the string to lowercase and remove spaces
    s = strings.ToLower(s)
    s = strings.ReplaceAll(s, " ", "")

    // Initialize two pointers
    left, right := 0, len(s)-1

    // Check characters from both ends moving towards the center
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func main() {
    var input string
    fmt.Println("Enter a string:")
    fmt.Scanln(&input)

    if isPalindrome(input) {
        fmt.Println("The string is a palindrome.")
    } else {
        fmt.Println("The string is not a palindrome.")
    }
}
