package array

import "strings"

func Palindrome(str string) bool {

	toMap := make(map[string]int)
	odd := 0
	for _, s := range str {

		toMap[strings.ToLower(string(s))]++
	}

	for _, i := range toMap {

		if odd > 1 {
			return false
		} else if i%2 != 0 {
			odd++
		}

	}

	return true
}
