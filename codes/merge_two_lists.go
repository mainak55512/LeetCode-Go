/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/

package codes

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var temp *ListNode = nil
	if list1 == nil && list2 != nil {
		return list2
	} else if list2 == nil && list1 != nil {
		return list1
	} else if list1 == nil && list2 == nil {
		return nil
	}
	if list1.Val < list2.Val {
		temp = &ListNode{
			Val:  list1.Val,
			Next: nil,
		}
		list1 = list1.Next
	} else {
		temp = &ListNode{
			Val:  list2.Val,
			Next: nil,
		}
		list2 = list2.Next
	}
	temp.Next = MergeTwoLists(list1, list2)
	return temp
}

func printList(listNode *ListNode) {
	if listNode.Next == nil {
		fmt.Println(listNode.Val)
		return
	}
	fmt.Println(listNode.Val)
	printList(listNode.Next)
}
