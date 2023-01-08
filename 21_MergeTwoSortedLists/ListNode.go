package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert(root *ListNode, elem int) *ListNode {
	temp := ListNode{Val: elem, Next: nil}

	if root == nil {
		root = &temp
		return root
	}

	curr := root
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &temp
	return root
}

func convertToList(arr []int) *ListNode {
	var head *ListNode
	for _, val := range arr {
		head = insert(head, val)
	}
	return head
}
