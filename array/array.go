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
