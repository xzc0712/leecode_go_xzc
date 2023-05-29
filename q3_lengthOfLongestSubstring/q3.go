package main

import (
	"fmt"
	"strings"
)

/*
思路（同向双指针，滑动窗口）
1. 遍历s，右端点向右移动，用map来记录字符的次数（初始为0）
2. 如果右端加入的字符在左侧子串内，则count+
3. 若count[]大于1，说明存在相同子串，则移动左端点，直到使count[]==1
4. max与right-left+1做比较，得出最长的子序列
*/

func maxnum(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	res := 0
	var count map[byte]int
	//range str类型返回的是十六进制的字符编码。例如：key:0 value:0x68
	//count := [128]int{}
	for right, c := range s {
		//left := right
		count[byte(c)]++
		for count[byte(c)] > 1 {
			count[s[left]]--
			left++
			//count[byte(c)]-- //注意是左侧的元素去掉，直到
		}
		res = maxnum(res, right-left+1)
	}
	return res
}

//func lengthOfLongestSubstring(s string) int {
//	left := 0
//	res := 0
//	//var count map[byte]int
//	count := [128]int{}
//	for right, c := range s {
//		//left := right
//		count[c]++
//		for count[c] > 1 {
//			count[s[left]]--
//			left++
//			//count[byte(c)]-- //注意是左侧的元素去掉，直到
//		}
//		res = maxnum(res, right-left+1)
//	}
//	return res
//}

/*
暴力思路（无法通过所有测试用例）
1. 从第一个字符开始遍历，如果子串中不包含当前字符，则加入，记录此时的最长长度。
2. 往后直到整个遍历结束
*/
func lengthOfLongestSubstring_xzc(s string) int {
	var s1 string
	max := 0
	temp := 0
	for i := 0; i < len(s)-1; i++ {
		s1 = string(s[i])
		for j := i + 1; j < len(s); j++ {
			index := strings.Contains(s1, string(s[j]))
			if !index {
				s1 = s1 + string(s[j])
			} else {
				temp = len(s1)
				break
			}
		}
		if temp > max {
			max = temp
		}
	}
	return max
}

func main() {
	s := "abcabcbb  "
	fmt.Printf("%d", lengthOfLongestSubstring(s))
}
