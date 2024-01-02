package main

import (
	"fmt"
	"math"
)

func maxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxEndingHere := nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = int(math.Max(float64(nums[i]), float64(maxEndingHere+nums[i])))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}

	return maxSoFar
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubarraySum(nums)
	fmt.Println(result) // Output: 6
}
