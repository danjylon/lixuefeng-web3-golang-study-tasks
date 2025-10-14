package main

/**
回文数
考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
import "fmt"

func main() {
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Println("error:", err)
	}
	// 回文数都是数字，不用担心len取字节数问题
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			fmt.Println("不是回文数")
			return
		}
	}
	fmt.Println("是回文数")
}
