package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev

		prev = slow
		slow = next
	}

	ans := 0
	for head != nil && prev != nil {
		ans = max(ans, head.Val+prev.Val)
		head = head.Next
		prev = prev.Next
	}

	return ans
}
