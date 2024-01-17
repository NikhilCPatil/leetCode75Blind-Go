package main

import (
		"fmt"
)


func wordBreak(s string, wordDist []string)bool{
d := make([]bool, len(s)+1) 
d[0] = true
wordMap := make(map[string]bool)
for _,w := range wordDist {
	wordMap[w] = true
}
// fmt.Println(wordMap)

for i := 1; i <= len(s); i++ {
	for j := 0; j < i; j++ {
		if d[j] && wordMap[s[j:i]]{
			d[i] = true
			break
		}
	}
}

// fmt.Println(d)
return d[len(s)]
}

func main(){
	s := "leetcode"
	wordDist := []string{"leet", "code"}
	result := wordBreak(s, wordDist)
	fmt.Println(result)
}