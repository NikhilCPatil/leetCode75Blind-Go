package main

import "fmt"

func numDecodings(s string) int {
	n:= len(s)
   
	  d := make([]int, n+1)
	   d[0] = 1
	  fmt.Println(d)
	   for i:= 1; i <= n; i++ {
		   if s[i-1] != '0' {
			  d[i] += d[i-1] 
		   }
  
		   if i > 1 && isDoubleNumberValid(s[i-2],s[i-1]) {
			   d[i] += d[i-2]
		   }
	   }
	   fmt.Println(d)
	   return d[n]
   }
   
   func isDoubleNumberValid(a , b byte)bool{
  if a == '1' {
	  return true
  }else if a == '2' && b <= '6' {
	  return true
  }
	  return false
   }

 func main(){
	nums := "06"
	result := numDecodings(nums)
	fmt.Println(result)
 }