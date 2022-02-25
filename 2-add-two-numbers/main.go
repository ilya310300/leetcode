package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var output *ListNode

	if l1 == nil && l2 == nil {
		return nil
	}

	var v1 int
	if l1 != nil {
		v1 = l1.Val
	}

	var v2 int
	if l2 != nil {
		v2 = l2.Val
	}

	s := v1 + v2

	q := s / 10
	r := s % 10

	var next *ListNode
	if q > 0 {
		next = &ListNode{
			Val:  q,
			Next: nil,
		}
	}

	if l1 != nil || l2 != nil {
		if l1 == nil {
			next = addTwoNumbers(next, l2.Next)
		} else if l2 == nil {
			next = addTwoNumbers(next, l1.Next)
		} else {
			next = addTwoNumbers(addTwoNumbers(l1.Next, l2.Next), next)
		}
	}

	output = &ListNode{
		Val:  r,
		Next: next,
	}

	return output
}
