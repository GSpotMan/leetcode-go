package algorithms

import "strings"

func getNext(s string, next []int) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j

	}
}
func strStr(haystack string, needle string) int {
	j := 0
	next := make([]int, len(needle))
	getNext(needle, next)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}

	}
	return -1
}

// 459 重复的子字符串
func repeatedSubstringPattern(s string) bool {
	len := len(s)
	ss := s + s
	if strings.Index(ss[1:], s) != -1 && strings.Index(ss[1:], s) != len {
		return true
	}
	return false
}

// 459 重复的子字符串 法2
func repeatedSubstringPattern2(s string) bool {
	len := len(s)
	next := make([]int, len)
	getNext(s, next)
	if next[len-1] > 0 && len%(len-next[len-1]) == 0 {
		return true
	}
	return false
}

// 二叉树前序遍历
func preorderTraversal(root *TreeNode) []int {

}
