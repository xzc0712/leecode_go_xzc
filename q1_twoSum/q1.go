package main

import "fmt"

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\n", a[i])
	}
}

func addNumArray(x int, y int) []int {
	var res []int
	res = make([]int, 0, 6)
	for i := 0; i < 5; i++ {
		res = append(res, x, y)
	}
	return res
}

func twoSum(nums []int, target int) []int {
	//var res [99]int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				//fmt.Printf("[%d,%d]", i, j)
				//res[0] = i
				//res[1] = j
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	n := 4
	//var n int
	target := 9
	//var arr = [...]int{2,7,11,15}
	var nums = make([]int, n, 10)
	//fmt.Scanln()
	for i := 0; i < n; i++ {
		fmt.Printf("请输入第%d个元素的值", i+1)
		fmt.Scanln(&nums[i])
	}
	//fmt.Scanln(&target)
	twoSum(nums, target)
	//fmt.Printf("%d", target)
	//printArray(nums)
}
