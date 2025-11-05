package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next   // 0->2
		next := head.Next.Next // 3: 下一对的第一个节点，需要预先保存
		head.Next.Next = head  // 2->1
		head.Next = next       // 1->3
		pre = head             // 第一次循环依赖dummy节点，第二次之后只需要指向节点对的第二个节点即可
		head = next            // 第二次循环之后头需要移动到下一个节点对的第一个节点
	}
	return dummy.Next
}
