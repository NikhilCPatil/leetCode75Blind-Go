package main

import (
	"fmt"
	"strings"
	"sort"
)

func validAnagram(strs []string)[][]string{
	m := make(map[string][]string)
	var res [][]string
	for _,word := range strs {
		m[sortStr(word)] =  append(m[sortStr(word)], word)
	}
	for _,v := range m {
		res = append(res, v)
	}
	return res
}

func sortStr(s string)string{
	res := strings.Split(s, "")
	sort.Strings(res)
	return strings.Join(res, "")
}

func main(){
	s :=  []string{"cat", "tac", "dog", "odg"}
	result := validAnagram(s)
	fmt.Println("Group Anagram :", result)
}