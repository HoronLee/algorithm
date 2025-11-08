package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var pA, pB = headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

// 双循环法
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	curA := headA
// 	curB := headB
// 	var lenA, lenB int
// 	for curA != nil {
// 		lenA++
// 		curA = curA.Next
// 	}
// 	for curB != nil {
// 		lenB++
// 		curB = curB.Next
// 	}
// 	curA = headA
// 	curB = headB
// 	if lenB > lenA {
// 		lenA, lenB = lenB, lenA
// 		curA, curB = curB, curA
// 	}
// 	gap := lenA - lenB
// 	for gap > 0 {
// 		curA = curA.Next
// 		gap--
// 	}
// 	for curA != nil {
// 		if curA == curB {
// 			return curA
// 		}
// 		curA = curA.Next
// 		curB = curB.Next
// 	}
// 	return nil
// }
