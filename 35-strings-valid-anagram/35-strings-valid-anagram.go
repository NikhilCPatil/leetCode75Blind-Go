package main

import (
	"fmt"
	"strings"
	"sort"
)

func validAnagram(s string, t string)bool{
	if len(s) != len(t){
		return false
	}

	m := make(map[byte]int)

	for i := range s {
		m[s[i]]++
		m[t[i]]--
	}
	
	for _, j := range m {
		if j != 0 {
			return false
		}
	}
	return true
}

func reveresStr(s string)string{
	res := ""
	for i := len(s)-1; i >= 0; i-- {
		res = res + string(s[i])
	}
	fmt.Println(res)
	return res
}

func sortStr(s string)string{
	res := strings.Split(s, "")
	sort.Strings(res)
	return strings.Join(res, "")
}

func main(){
	s := "cat"
	t := "cat"
	result := validAnagram(s,t)
	fmt.Println("Valid Anagram? :", result)
}