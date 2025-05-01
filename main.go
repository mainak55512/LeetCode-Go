package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	// fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func printList(listNode *ListNode) {
	if listNode.Next == nil {
		fmt.Println(listNode.Val)
		return
	}
	fmt.Println(listNode.Val)
	printList(listNode.Next)
}
