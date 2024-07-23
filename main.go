package main

import (
	"cracking_the_coding/array"
	"fmt"
	"strings"
)

func main() {

	//var test bool = checkReturnWord("pale", "bake")
	//
	//fmt.Printf("%v \n", test)
	//fmt.Print("finish \n")
	res := array.IsUniqueChars("abc\n")
	fmt.Printf("res : %v", res)
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
