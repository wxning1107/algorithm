package algorithm

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	curr := res
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: (n1 + n2 + carry) % 10, Next: nil}
		carry = (n1 + n2 + carry) / 10
		curr = curr.Next
	}

	return res.Next
}
