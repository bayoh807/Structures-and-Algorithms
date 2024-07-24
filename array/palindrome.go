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

func Palindrome1(str string) bool {

	toMap := make(map[string]int)
	odd := 0
	for _, s := range str {
		lowerStr := strings.ToLower(string(s))
		if lowerStr != " " {
			toMap[lowerStr]++
		}

		if val, _ := toMap[lowerStr]; val%2 == 1 {
			odd++
		} else {
			odd--
		}

	}

	return odd <= 1
}
