package main

import "fmt"

/**
两数之和
考察：数组遍历、map使用
题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
链接：https://leetcode-cn.com/problems/two-sum/
*/

func main() {
	//nums := []int{2, 7, 11, 15}
	//target := 9
	nums := []int{3, 2, 4}
	target := 6
	sum := twoSum(nums, target)
	fmt.Println(sum)
}
func twoSum(nums []int, target int) []int {
	mp := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			//t := nums[i] + nums[j]
			//fmt.Printf("i:%d,j:%d,t:%d, %d+%d=%d\n", i, j, t, nums[i], nums[j], t)
			_, ok := mp[nums[i]+nums[j]]
			if !ok {
				mp[nums[i]+nums[j]] = []int{i, j}
			}
		}
	}
	return mp[target]
}
