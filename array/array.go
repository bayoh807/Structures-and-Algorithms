package array

import (
	"sort"
	"strings"
)

func IsUniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	} else {
		var charSet [128]bool

		for _, s := range str {

			if val := charSet[s]; val {
				return false
			} else {
				charSet[s] = true
			}

		}
		return true
	}
}

func Permutation(str1, str2 string) bool {

	sortString := func(s string) string {
		str := strings.Split(s, "")
		sort.Strings(str)
		return strings.Join(str, "")
	}

	if len(str1) != len(str2) || sortString(str1) != sortString(str2) {
		return false
	} else {
		return true
	}
}

func Permutation2(str1, str2 string) bool {

	var sliceArr [128]int

	if count := len(str1); count != len(str2) {
		return false
	} else {
		for _, s := range str1 {
			sliceArr[s]++
		}

		for _, s := range str2 {
			if check := sliceArr[s] - 1; check < 0 {
				return false
			}
		}
		return true
	}

}
