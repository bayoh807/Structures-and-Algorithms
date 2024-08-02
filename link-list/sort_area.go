package link_list

func AddList(l1, l2 *LinkedListNode) *LinkedListNode {

	var head, item *LinkedListNode
	var forward int

	for l1 != nil || l2 != nil {
		total := l1.Value + l2.Value + forward
		current := &LinkedListNode{
			Value: (total) % 10,
		}

		forward = total / 10

		if head == nil {
			head = current
		} else {
			item.Next = current
		}
		item = current
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = &LinkedListNode{}
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = &LinkedListNode{}
		}
	}

	return head

}
