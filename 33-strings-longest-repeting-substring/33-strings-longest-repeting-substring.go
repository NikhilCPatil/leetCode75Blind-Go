package main

import "fmt"

func longestReapetedSub(s string, k int)int{
	charMap := make(map[byte]int)
	maxLen := 0
	start := 0
	maxFreq := 0

	for end := 0; end < len(s); end++ {
		charMap[s[end]]++
		maxFreq = max(maxFreq, charMap[s[end]])
		if end-start+1-maxFreq  > k{
			charMap[s[start]]--
			start++
		}

		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

func max(a int, b int)int{
	if a > b {
		return a
	}
	return b
}

func main(){
	k:= 1
	s := "ABABCCCB"
	result := longestReapetedSub(s, k)
	fmt.Println(result)
}