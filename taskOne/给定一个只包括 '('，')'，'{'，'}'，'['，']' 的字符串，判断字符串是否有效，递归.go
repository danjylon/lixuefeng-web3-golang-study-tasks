package main

/**
有效的括号
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
链接：https://leetcode-cn.com/problems/valid-parentheses/
*/
import (
	"fmt"
	"strings"
)

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(input)
	flag := isValid(input)
	fmt.Println("是否有效: ", flag)
}

func isValid(s string) bool {
	mp := map[string]string{"(": ")", ")": "(", "[": "]", "]": "[", "{": "}", "}": "{"}
	m := make(map[string]int)
	for _, ch := range s {
		m[string(ch)]++
	}
	fmt.Println(m)
	flag := recursion(s, mp, m)
	return flag
}

// 递归
func recursion(input string, mp map[string]string, m map[string]int) bool {
	if len(input) == 0 {
		return true
	}
	for _, sOneSide := range input {
		sOtherSide := mp[string(sOneSide)]
		fmt.Println(m[string(sOneSide)], m[sOtherSide])
		if m[string(sOneSide)] != m[sOtherSide] {
			fmt.Println(false)
			return false
		}
	}
	for i := 0; i < len(input); i++ {
		sOneSide := string(input[i])
		if sOneSide == ")" || sOneSide == "]" || sOneSide == "}" {
			continue
		}
		fmt.Println("s: ", sOneSide)
		sOtherSide := mp[sOneSide]
		j := strings.Index(input, sOtherSide)
		if j > i {
			input = input[i+1 : j]
			if len(input)%2 != 0 {
				fmt.Println("len(input) %2 !=0: ", false)
				return false
			}
			m := make(map[string]int)
			recursion(input, mp, m)
		}
	}
	return true
}
