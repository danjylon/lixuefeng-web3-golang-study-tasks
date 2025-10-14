package main

/**
最长公共前缀
考察：字符串处理、循环嵌套
题目：查找字符串数组中的最长公共前缀
链接：https://leetcode-cn.com/problems/longest-common-prefix/
*/
import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)
	fmt.Println("prefix:", prefix)
}
func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for _, s := range strs {
		if minLen > len(s) {
			minLen = len(s)
		}
	}
	prefix := ""
loop:
	for i := 0; i < minLen; i++ {
		letter := strs[0][i]
		for _, s := range strs {
			if letter != s[i] {
				break loop
			}
		}
		prefix = prefix + string(letter)
	}
	return prefix
}
