package main

import "fmt"

/*
*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("nums:", nums)
	sliceMultiplyTwo(&nums)
	fmt.Println("nums:", nums)
	sliceMultiplyTwo2(nums)
	fmt.Println("nums:", nums)
}
func sliceMultiplyTwo(arr *[]int) {
	for i, num := range *arr {
		(*arr)[i] = num * 2
	}
}
func sliceMultiplyTwo2(arr []int) {
	for i, num := range arr {
		arr[i] = num * 2
	}
}
