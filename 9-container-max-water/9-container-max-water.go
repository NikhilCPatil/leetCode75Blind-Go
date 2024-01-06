package main

import "fmt"

//complexity 0[n]

func calculateMaxWaterArea(height []int)int{
	maxWater := 0
	left, right := 0, len(height)-1

    for left < right {
        h := min(height[left], height[right])
        w := right - left
        a := h * w
        maxWater = max(a, maxWater)
        if height[left] < height[right]{
            left++
        }else {
            right--
        }
    }

    return maxWater
	
}

func max(a int,b int)int{
    if a > b {
        return a
    }
return b
}

func min(a int,b int)int{
    if a < b {
        return a
    }
return b
}

func main(){
heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
result := calculateMaxWaterArea(heights)
fmt.Println(result)
}