package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

/**
* Find Middel
* Reverse Second half
* Merge first and second half
*/

func reorderList(head *ListNode){
	if(head == nil || head.Next == nil){
		return 
	}
 mid := findMiddle(head)
//  printList(mid)
 rev := reverseList(mid.Next)
//  printList(rev)
 mid.Next = nil
mergeLists(head, rev)
}

func reverseList(head *ListNode)*ListNode{
	var next, prev *ListNode
	current := head
	for current != nil{
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func findMiddle(head *ListNode) *ListNode{
	slow, fast := head, head
	for  fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeLists(l1 *ListNode, l2 *ListNode){
	for l2 != nil {
		next1, next2 := l1.Next, l2.Next
		l1.Next, l2.Next = l2, next1
		l1,l2 = next1, next2 
	}
}

func printList(head *ListNode){
current := head
for current != nil{
	fmt.Printf("%d ->", current.Val)
	current = current.Next
}
fmt.Printf("\n")
}

func main(){
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	printList(list)
	reorderList(list)
	printList(list)
}