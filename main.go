package main

import (
	"CPGo/codes"
	"fmt"
)

func main() {
	// fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	// fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(codes.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
