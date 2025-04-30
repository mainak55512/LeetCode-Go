package main

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, elem := range strs[1:] {
		i := 0
		for ; i < len(prefix) && i < len(elem) && prefix[i] == elem[i]; i++ {
		}
		prefix = prefix[:i]
	}
	return prefix
}
