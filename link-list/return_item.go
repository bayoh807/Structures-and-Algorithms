package link_list

func ReturnValue(list *LinkedListNode, node int) *LinkedListNode {

	var count int
	hashTable := make(map[int]*LinkedListNode)
	for {
		hashTable[count] = list
		if list.Next != nil {
			list = list.Next
			count++
		} else {
			break
		}
	}

	return hashTable[count-node+1]
}

func ReturnValue2(list *LinkedListNode, node int) *LinkedListNode {

	p1, p2 := list, list

	for i := 0; i < node; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
