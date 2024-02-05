package main

import (
	"fmt"
	"math"
)

func minimumWindow(s string, t string)string{
	charMap := make(map[byte]int)
	right, left, requiredCount := 0, 0, len(t)
	minWindow := ""
	minLen := math.MaxInt32

	for i := 0; i < len(t); i++ {
		charMap[t[i]]++
	}
	for right < len(s) {
		if charMap[s[right]] > 0{
			requiredCount--
		}

		charMap[s[right]]--
		right++
		for requiredCount == 0 {
			
			if(right - left < minLen){
				minLen = right - left
				minWindow = s[left:right] 
			}

			charMap[s[left]]++
			if charMap[s[left]] > 0 {
				requiredCount++
			}
			left++
		}	
	}
	return minWindow
}

func main(){
	s := "AJBCOINABC"
	t := "ABC"
	result := minimumWindow(s, t)
	fmt.Println(result) 
}