package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	var slow, fast = dummyHead, dummyHead

	// 1. 建立 "n+1" 的间距 - 步骤 A：
	//    让 fast 指针先走 n 步。
	//    (循环条件 fast.Next != nil 保证了 fast 不会越过链表末尾)
	for fast.Next != nil && n > 0 {
		fast = fast.Next
		n--
	}

	// 2. 建立 "n+1" 的间距 - 步骤 B：
	//    fast 再多走一步。
	//    *关键点*：执行完这一步后，fast 严格领先 slow (n+1) 步。
	//    *为什么是 n+1？*
	//    因为我们的目标是让 slow 停在 *待删除节点的前一个节点* 上。
	//    当 fast 走到 nil (链表末尾的再下一个) 时，
	//    slow 刚好会停在 (倒数第 n+1) 个节点上，即倒数第n个节点的前驱。
	fast = fast.Next

	// 3. 快慢指针同时推进：
	//    现在 fast 和 slow 之间保持着 (n+1) 的恒定间距。
	//    它们一起移动，直到 fast 遇到 nil。
	for fast != nil {
		fast = fast.Next // fast 先走
		slow = slow.Next // slow 跟上
	}

	// 4. 执行删除：
	//    当 fast 变为 nil 时，slow 正好位于倒数第 n 个节点的 *前一个* 节点。
	//    (例如：删除倒数第2个(4)，slow 停在倒数第3个(3)上)
	//    通过修改 "slow.Next"，我们让 slow 直接指向 "slow.Next.Next"，
	//    从而跳过了（即删除了）目标节点。
	slow.Next = slow.Next.Next // 直接断开第n个元素的链接

	// 5. 返回新链表的头节点：
	//    必须返回 dummyHead.Next，因为 head 节点本身可能已经被删除了。
	return dummyHead.Next
}
