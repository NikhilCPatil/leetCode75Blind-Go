package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) }
func (h ListNodeHeap) Less(i,j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i,j int) { h[i],h[j] = h[j], h[i] }

func (h *ListNodeHeap)Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap)Pop()interface{} {
	old := *h 
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode)*ListNode{
	pq := make(ListNodeHeap, 0)
	heap.Init(&pq)

	for _, list := range lists {
		if list != nil {
			heap.Push(&pq, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	for pq.Len() > 0 {
		
		node := heap.Pop(&pq).(*ListNode)
		current.Next = node
		current = current.Next

		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}

	return dummy.Next
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main(){
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}


	mergedList := mergeKLists([]*ListNode{list1, list2, list3})
	printList(mergedList)
}