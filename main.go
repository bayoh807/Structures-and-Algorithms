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
		Value: 4,
	}
	Two := &link_list.LinkedListNode{
		Value: 1,
	}

	Three := &link_list.LinkedListNode{
		Value: 6,
	}
	Four := &link_list.LinkedListNode{
		Value: 2,
	}

	//Five := &link_list.LinkedListNode{
	//	Value: 1,
	//}

	One.Next = Two
	Two.Next = Three
	Three.Next = Four
	Four.Next = Three

	//list := []*link_list.LinkedListNode{
	//	One, Two, Three, Four, Five,
	//}
	res := link_list.DeleteRepeat2(One)
	//list := One

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
