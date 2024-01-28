package main

import "fmt"

 type Node struct {
	data int
	next *Node
}

func printLinkedList(head *Node){
	current := head
	for current != nil {
		fmt.Printf("%d ->",current.data)
		current = current.next
	}
}

func removenthinkedlist(list1 *Node, n int)*Node{
	dummy := &Node{next: list1}
	fast, slow := dummy, dummy

	for i := 0; i <= n; i++ {
		fast = fast.next
	}

	for fast != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next

	return dummy.next
}

func main(){
	list1 := &Node{ data : 1, next : &Node{ data: 2, next : &Node{data: 3, next: &Node{data: 5, next: nil}} } }
	newList := removenthinkedlist(list1, 2)

	printLinkedList(newList)
}