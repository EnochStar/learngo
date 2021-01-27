package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//	var cur *ListNode
//	for head.Next != nil {
//		next := head.Next
//		head.Next = cur
//		cur = head
//		head = next
//	}
//	return cur
//}
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
