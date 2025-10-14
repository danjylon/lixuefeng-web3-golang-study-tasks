package main

/**
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
import "fmt"

func main() {
	//intervals := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	//intervals := [][]int{[]int{3, 4}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	//intervals := [][]int{[]int{2, 6}, []int{1, 3}, []int{8, 10}, []int{15, 18}}
	//intervals := [][]int{[]int{2, 6}, []int{3, 4}, []int{8, 10}, []int{15, 18}}
	intervals := [][]int{[]int{2, 6}, []int{8, 10}, []int{15, 18}, []int{3, 4}}
	//intervals := [][]int{[]int{1, 4}, []int{4, 5}}
	//intervals := [][]int{[]int{4, 7}, []int{1, 4}}
	sort(intervals)
	fmt.Println("intervals:", intervals)
	fmt.Println(merge(intervals))
}
func merge(intervals [][]int) [][]int {

	var result [][]int
	for i := 0; i < len(intervals); i++ {
		iInts := intervals[i]
		temp := iInts[:]
		for j := i + 1; j < len(intervals); j++ {
			jInts := intervals[j]
			//只要一个的iInts[1]大于等于jInts[0]，且iInts[0]<=大于等于jInts[1]
			if iInts[1] >= jInts[0] && iInts[0] <= jInts[1] {
				//fmt.Println("iInts:", iInts)
				//fmt.Println("jInts:", jInts)
				temp = findMinMax(iInts, jInts)
				fmt.Println("temp:", temp)
				fmt.Println("iInts:", iInts)
				i++
			}
		}
		result = append(result, temp)
	}
	return result
}

func sort(intervals [][]int) {
	//var result [][]int
	for i := 0; i < len(intervals); i++ {
		iInts := intervals[i]
		for j := i + 1; j < len(intervals); j++ {
			jInts := intervals[j]
			if iInts[0] >= jInts[0] {
				temp := jInts
				intervals[j] = iInts
				intervals[i] = temp
			}
		}
	}
}

func findMinMax(iInts, jInts []int) []int {
	nums := append(iInts, jInts...)
	if len(nums) == 0 {
		return []int{0, 0}
	}
	minInt, maxInt := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minInt {
			minInt = num
		}
		if num > maxInt {
			maxInt = num
		}
	}
	return []int{minInt, maxInt}
}
