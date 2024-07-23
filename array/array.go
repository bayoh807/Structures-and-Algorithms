package array

import "strings"

func IsUniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	} else {
		charSet := map[string]bool{}
		strSplit := strings.Split(str, "")
		for _, s := range strSplit {

			if val := charSet[s]; val {
				return false
			} else {
				charSet[s] = true
			}

		}
		return true
	}

}

func Permutation(str, t string) bool {

}
