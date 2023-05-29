package main

/*
q4:寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。
请你找出并返回这两个正序数组的中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。
。*/
import "fmt"

/**思路:
1. 首先需要判断奇数偶数的情况;取巧的办法是计算（m+n+1）/2和(m+n+2)/2
2.

*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64

	return res
}

func main() {
	//a := [1000]int{1, 3}
	//b := [1000]int{2}
	var n, m int
	var a = make([]int, n, 10)
	var b = make([]int, m, 10)
	for i := 0; i < n; i++ {
		fmt.Printf("请输入数组a第%d个元素的值", i+1)
		fmt.Scanln(&a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Printf("请输入数组b第%d个元素的值", i+1)
		fmt.Scanln(&b[i])
	}
	fmt.Printf("%.5f", findMedianSortedArrays(a, b))
}
