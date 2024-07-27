package main

import (
	"cracking_the_coding/array"
	"fmt"
	"strings"
)

func main() {

	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	res := array.Rotate(arr)
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
