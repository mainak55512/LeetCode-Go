/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.
*/

package codes

func LongestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, elem := range strs[1:] {
		i := 0
		for ; i < len(prefix) && i < len(elem) && prefix[i] == elem[i]; i++ {
		}
		prefix = prefix[:i]
	}
	return prefix
}
