package main

import "fmt"

func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    // Initialize result array with product of elements to the left of each index
    leftProduct := 1
    for i := 0; i < n; i++ {
        result[i] = leftProduct
        leftProduct *= nums[i]
    }
	// result = [ 1, 1, 2, 6]

	fmt.Println(result)

    // Update result array with product of elements to the right of each index
    rightProduct := 1
    for i := n - 1; i >= 0; i-- {
        result[i] *= rightProduct
		rightProduct *= nums[i]
		fmt.Println(rightProduct)
	}
	// result = [ 24, 12, 8, 6]

	fmt.Println(result)

    return result
}

func main() {
    nums := []int{1, 2, 3, 4}
    result := productExceptSelf(nums)
    fmt.Println(result) // Output: [24 12 8 6]
}
