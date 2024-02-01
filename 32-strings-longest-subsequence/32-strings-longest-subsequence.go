package main

import "fmt"

func longestSubSeq(s string)int{
	charMap := make(map[byte]int)
	start := 0
	maxLen := 0

	for end := 0; end < len(s); end++ {
		if index, found := charMap[s[end]]; found {
			start = max(start , index+1)
		}

		charMap[s[end]] = end

		maxLen = max(maxLen, end-start+1)
		
	}
	return maxLen
}

func max(a int,b int)int{
if a > b {
	return a
}
return b
}

func main(){
 s := "abcabcbbcc"
 result := longestSubSeq(s) 
 fmt.Println(result)
}