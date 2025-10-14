package main

/**
删除有序数组中的重复项：
给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
*/
import "fmt"

func main() {
	//nums := []int{1, 1, 2}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{0, 1, 2, 3, 4}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)
}
func removeDuplicates(nums []int) int {
	num := 1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				continue
			} else {
				nums[num] = nums[j]
				num += 1
				i = j - 1
				//i++
				break
			}
		}
	}
	fmt.Println("nums:", nums)
	return num
}
