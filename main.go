package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node3 := ListNode{
		Val:  4,
		Next: nil,
	}
	node2 := ListNode{
		Val:  2,
		Next: &node3,
	}
	node1 := ListNode{
		Val:  1,
		Next: &node2,
	}
	list1 := &node1
	node_n3 := ListNode{
		Val:  4,
		Next: nil,
	}
	node_n2 := ListNode{
		Val:  3,
		Next: &node_n3,
	}
	node_n1 := ListNode{
		Val:  1,
		Next: &node_n2,
	}
	list2 := &node_n1
	merged := mergeTwoLists(list1, list2)

	printList(merged)
}

func printList(listNode *ListNode) {
	if listNode.Next == nil {
		fmt.Println(listNode.Val)
		return
	}
	fmt.Println(listNode.Val)
	printList(listNode.Next)
}
