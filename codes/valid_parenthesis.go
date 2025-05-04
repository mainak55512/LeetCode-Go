/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true



Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package codes

func IsValid(s string) bool {
	var parenStack []rune
	for _, elem := range s {
		if elem == '(' || elem == '{' || elem == '[' {
			parenStack = append(parenStack, elem)
		} else if len(parenStack) > 0 {
			switch elem {
			case ')':
				{
					if parenStack[len(parenStack)-1] != '(' {
						return false
					}
				}
			case '}':
				{
					if parenStack[len(parenStack)-1] != '{' {
						return false
					}
				}
			case ']':
				{
					if parenStack[len(parenStack)-1] != '[' {
						return false
					}
				}
			}
			parenStack = parenStack[:len(parenStack)-1]
		} else {
			return false
		}
	}
	if len(parenStack) > 0 {
		return false
	}
	return true
}
