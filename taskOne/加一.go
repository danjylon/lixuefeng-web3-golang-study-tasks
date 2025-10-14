package main

/**
加一
考察：数组操作、进位处理
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
链接：https://leetcode-cn.com/problems/plus-one/
*/
import (
	"fmt"
	"math"
)

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(digits)
	res := plusOne(digits)
	fmt.Println("result:", res)
}
func plusOne(digits []int) []int {
	num := 0
	for i := 0; i < len(digits); i++ {
		num += digits[i] * int(math.Pow10(len(digits)-i-1))
	}
	num += 1
	var res []int
	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}
	Reverse(res)
	return res
}

// Reverse reverses a slice in place
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
