package main

import (
	link_list "cracking_the_coding/link-list"
	"fmt"
	"strings"
)

func main() {

	//arr := [][]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//	{13, 14, 15, 16},
	//}
	One := &link_list.LinkedListNode{
		Value: 7,
	}
	Two := &link_list.LinkedListNode{
		Value: 1,
	}

	Three := &link_list.LinkedListNode{
		Value: 6,
	}
	Four := &link_list.LinkedListNode{
		Value: 5,
	}

	Five := &link_list.LinkedListNode{
		Value: 9,
	}
	Six := &link_list.LinkedListNode{
		Value: 2,
	}

	One.Next = Two
	Two.Next = Three
	Four.Next = Five
	Five.Next = Six

	res := link_list.AddList(One, Four)

	fmt.Printf("res : %v\n", res)
}

func checkReturnWord(input, compare string) bool {

	toArrInput := stringToArr(input)
	toArrCompare := stringToArr(compare)
	//toCheckMap := map[string]int{}
	res := len(toArrInput) - len(toArrCompare)

	switch true {
	case res == -1:
	case res == 1, res == 0:
	default:
		return false
	}
	//if res := len(toArrInput) - len(toArrCompare); res > 1 || res < -1 {
	//	return false
	//}

	//for s,i := range toArrInput
	return true
}

func stringToArr(input string) []string {
	return strings.Split(input, "")
}
