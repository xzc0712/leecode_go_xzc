package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路：
1.判断两个链表长度
2.短的补零
3.末尾相加，如果大于0，则取余数，进位为1；小于零，则直接取，进位为0。
4.直到两个链表遍历结束
5.最后查看是否有进位，如果有，则最后位补1。
*/

//func getListNodeLen(l *ListNode) int {
//	var len int
//	var p *ListNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var res *ListNode
	c := 0
	//var isAdd bool
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil { //把值取出来，向后遍历
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + c
		sum, c = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum} // 新建一个节点，值为sum
			res = head                 //指向head
		} else {
			res.Next = &ListNode{Val: sum}
			res = res.Next
		}
	}
	if c > 0 {
		res.Next = &ListNode{Val: c}
	}
	return head
}

func main() {
	fmt.Println("xzc_q2")
}
