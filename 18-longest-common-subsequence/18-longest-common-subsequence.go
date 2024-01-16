package main

import (
		"fmt"
)


func longestCommonSubsequence(text1 string, text2 string)int{

m, n:= len(text1), len(text2)

d := make([][]int, m+1) 
for i := range d {
	d[i] = make([]int, n+1)
}

for i := 1; i <= m; i++ {
	for j := 1; j <= n; j++ {
		if text1[i-1] == text2[j-1]{
			d[i][j] = 1 + d[i-1][j-1] 
		}else{
			d[i][j] = max(d[i-1][j], d[i][j-1])
		}
	}
}


fmt.Println(d)
return d[m][n]
}

func max(a int, b int) int{
	if(a > b){
		return a
	}
	return b
}



func main(){
	text1 := "abcde"
	text2 := "ace"
	result := longestCommonSubsequence(text1, text2)
	fmt.Println(result)
}