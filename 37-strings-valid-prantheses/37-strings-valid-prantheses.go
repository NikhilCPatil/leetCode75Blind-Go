package main

import (
	"fmt"
)

func validParantheses(s string)bool{
	charMap :=  map[rune]rune{
		'}' : '{',
		')' : '(',
		']' : '['}
	
	stack := make([]rune, 0)
	
	for _, char := range s {
		if isOppning(char){
			stack = append(stack, char)
		}else{
			if len(stack) == 0 || stack[len(stack)-1] != charMap[char] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isOppning(s rune)bool{
	return s == '[' || s == '(' || s == '{'
}

func main(){
	s :=  "{[]}}"
	result := validParantheses(s)
	fmt.Println("Group Anagram :", result)
}