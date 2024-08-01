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
