package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//生成链表

func makeListNode(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	node := &ListNode{
		Val: s[0],
	}
	tmp := node
	for i := 1; i < len(s); i++ {
		tmp.Next = &ListNode{
			Val: s[i],
		}
		tmp = tmp.Next
	}
	return node

}

// 合并链表

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//如果有一条链是nil，直接返回另外一条链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 定义一个结果节点
	var res *ListNode
	// 当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}

	return res

}

// 遍历链表

func (l *ListNode) PrintListNodeASC() {
	var s []int
	//fmt.Println(l)
	ptr := l
	//fmt.Println(ptr)
	for ptr != nil {
		//fmt.Println(ptr.Val)
		s = append(s, ptr.Val)
		ptr = ptr.Next
	}
	fmt.Println(s)
}

func main() {
	var (
		l1 *ListNode
		l2 *ListNode
	)

	s1 := []int{0, 0, 4}
	s2 := []int{1, 3, 4}
	l1 = makeListNode(s1)
	//l1.PrintListNodeASC()
	l2 = makeListNode(s2)
	//l2.PrintListNodeASC()

	mergeTwoLists(l1, l2).PrintListNodeASC()

}
