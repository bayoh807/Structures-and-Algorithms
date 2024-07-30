package link_list

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

func DeleteRepeat(list *LinkedListNode) *LinkedListNode {

	toMap := make(map[LinkedListNode]bool)
	var previous *LinkedListNode

	for {
		if _, res := toMap[*list]; res {
			previous.Next = nil
			return list
		} else {
			toMap[*list] = true
			previous = list
		}

		list = list.Next
	}

}

func DeleteRepeat2(list *LinkedListNode) *LinkedListNode {

	current := list
	for current != nil && current.Next != nil {
		// Use runner to find and remove duplicates
		runner := current
		for runner.Next != nil {
			if runner.Next.Value == current.Value {
				// Remove the duplicate
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}

			// Check if we've looped back to current or its known neighbors
			if runner.Next == current || runner.Next == current.Next {
				runner.Next = nil // Break the cycle
				return list
			}
		}

		// Move to next unique element
		if current.Next != nil {
			current = current.Next
		} else {
			break
		}
	}

	return list

}
