package main

import "fmt"


func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    minPrice := prices[0]
    maxProfit := 0

    for i,num := range prices {
        if i == 0{
            continue
        }
            if num < minPrice {
                minPrice = num 
            }else {
                profit := num - minPrice
                if profit > maxProfit {
                    maxProfit = profit
                }
            }
    }
    return maxProfit
}


func main() {
    nums := []int{5,3,2,4}
	result := maxProfit(nums)
	
	fmt.Println("Max profit", result);

}
