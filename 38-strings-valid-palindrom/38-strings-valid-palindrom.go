package main

import (
	"fmt"
	"strings"
)

func validPalandrom(s string)bool{
	s = preProcess(s)
	left := 0
	right := len(s)-1

	for left < right{
		if s[left] != s[right]  {
			return false
		}
		left++
		right--
	}

	return true
}

func preProcess(s string)string{
	var result strings.Builder
	for _, char := range s {
		if isAlphanumeric(char) {
			result.WriteRune(toLower(char))
		}
	}
	return result.String()
}

func toLower(char rune)rune{
	if char >= 'A' && char <= 'Z' {
		return char + 32
	}
	return char
}

func isAlphanumeric(char rune)bool{
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}


func main(){
	s := "A man, a plan, a canal: Panama"
	fmt.Println("Valid Palandrom? ", validPalandrom(s))
}