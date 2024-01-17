package main

import (
		"fmt"
)


// func fibNormal(n int)int{
// 	if n == 1 || n== 2{
// 		return 1
// 	}else{
// 		return fibNormal(n-1) + fibNormal(n-2)
// 	}
// }

func fibBottomUp(n int)int{
	if n == 1 || n== 2{
		return 1
	}else{
		dp := make([]int, n + 1)
dp[1] = 1
dp[2] = 1

for i := 3; i <= n; i++ {
	dp[i] = dp[i-1] + dp[i-2]
}
return dp[n]
	}

}


func main(){
	n := 50
	result := fibBottomUp(n)
	fmt.Println(result)
}