package main

import (
	"fmt"
	"strings"
)

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
输入：strs = ["flower","flow","flight"]
输出："fl"
*/

/*思路：
1. 以第一个字符串作为基准元素，与后面的字符串进行比较。
2. 当后一个字符包含该基准元素时（string.index==0）则与下一个字符串进行比较
3. 当不包括基准元素时，则对基准元素进行末尾-1，直到包含为止
4. 基准元素为0，则返回空
ps：strings.Index()返回后一个字符在前一个字符中出现的位置
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	var res string
	res = strs[0]
	for _, k := range strs {
		for strings.Index(k, res) != 0 {
			res = res[:len(res)-1]
			if len(res) == 0 {
				return ""
			}
		}
	}
	return res
}
func main() {
	//str := make([]string,200)
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

}
