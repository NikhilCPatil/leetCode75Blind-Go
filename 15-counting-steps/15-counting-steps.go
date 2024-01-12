package main

import "fmt"

func climbsteps(n int) int {
    if n <= 2 {
        return n
    } 

    dp := make([]int, n + 1)
    dp[1] = 1
    dp[2] = 2

    for i:= 3; i <= n;i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

func main(){
    steps := 6
    result := climbsteps(steps)
    fmt.Println(result)
}