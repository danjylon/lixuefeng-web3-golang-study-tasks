package main

/**
有效的括号
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
链接：https://leetcode-cn.com/problems/valid-parentheses/
*/
import "fmt"

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(input)
	fmt.Println(isValidWithStack(input))

}
func isValidWithStack(s string) bool {
	mp := map[string]string{"(": ")", ")": "(", "[": "]", "]": "[", "{": "}", "}": "{"}
	stack := Stack{} // 初始化
	lastChar := ""
	for _, ch := range s {
		oneSide := string(ch)
		// 入栈（Push）
		stack.Push(oneSide)
		if oneSide == mp[lastChar] {
			// 出栈（Pop）
			pop1 := stack.Pop()
			fmt.Println("pop1:", pop1)
			// 出栈（Pop）
			pop2 := stack.Pop()
			fmt.Println("pop2:", pop2)
			// 查看栈顶（Peek）
			peek := stack.Peek()
			if peek != nil {
				lastChar = stack.Peek().(string)
			} else {
				lastChar = ""
			}
		} else {
			lastChar = oneSide
		}
	}
	return stack.IsEmpty()
}

type Stack []interface{}

// 入栈
func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

// 出栈
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

// 查看栈顶元素
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return (*s)[len(*s)-1]
}

// 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// 获取栈长度
func (s *Stack) Len() int {
	return len(*s)
}
