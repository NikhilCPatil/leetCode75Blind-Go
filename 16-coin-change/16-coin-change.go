package main

import (
		"fmt"
		"math"
)


func changeCalculater(coins []int ,amount int)int{

	d := make([]int, amount+1)

	for i := 1; i <=amount; i++ {
		d[i] = math.MaxInt32 
	}

	for _, coin:= range coins{
		for i := coin; i <= amount; i++{
			d[i] = min(d[i], d[i-coin]+1)
		}
	}

	if(d[amount] == math.MaxInt32){
		return -1
	}

	return d[amount]
}

func min(a int, b int) int{
	if(a < b){
		return a
	}
	return b
}

func main(){
	coins := []int{1,2,5}
	amount := 11
	result := changeCalculater(coins, amount)
	fmt.Println(result)
}