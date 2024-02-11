package main

import (
	"fmt"
)

func getPalindromicSubstring(s string) string{
	if(len(s) == 0){
		return ""
	}
	n := len(s)
	dp := make([][]bool, len(s))

	for i := range s {
		dp[i] = make([]bool, n)
	}

	
	// for 1 char
	for i := 0; i < n -1; i++ {
		dp[i][i] = true
	}

	maxLen := 1
	start := 0

	// for 2 char
	for i := 0; i < n -1; i++ {
		if(s[i] == s[i+1]){
			dp[i][i+1] = true
			maxLen = 2
			start = i
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i < n-k+1; i++ {
			j := i+k-1
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				start = i
				maxLen = k
			}
		}
	}

	return s[start : start+maxLen]
}

func main()  {
	s := "ccc"
	result := getPalindromicSubstring(s)
	fmt.Println("Palindromic Substring: ", result)
}