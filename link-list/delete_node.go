package link_list

func DeleteMiddleNode(list *LinkedListNode) *LinkedListNode {
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

	preMid := (count - 1) / 2

	preMidNode, _ := hashTable[preMid]
	midNode, _ := hashTable[preMid+1]
	preMidNode.Next = midNode.Next
	return list
}
