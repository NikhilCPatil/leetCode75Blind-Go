package main

import "fmt"

func isPartation(nums []int) bool {
    dp := make([][]bool, len(nums)+1)
    sum := 0
    for i := range nums {
        sum += nums[i]
    }

    if (sum & 1) == 1 {
        return false
    }

    sum /= 2
    for i := range dp {
        dp[i] = make([]bool, sum+1)
    }
    return checkPartation(nums, 0, sum, dp)
}

func checkPartation(nums []int, index int, sum int, dp [][]bool) bool {
    if sum < 0 || index == len(nums) {
        return false
    }
    if sum == 0 {
        return true
	}
	
	if dp[index][sum] {
		return dp[index][sum]
	}

    res := checkPartation(nums, index+1, sum, dp) || checkPartation(nums, index+1, sum-nums[index], dp)

    dp[index][sum] = res
    return res
}

func main() {
    nums := []int{1, 5, 5, 11}
    res := isPartation(nums)
    fmt.Println("Is Partation ?", res)
}