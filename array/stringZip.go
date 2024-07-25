package array

import (
	"sort"
	"strconv"
	"strings"
)

func sortString(s string) string {
	str := strings.Split(s, "")
	sort.Strings(str)
	return strings.Join(str, "")
}

func ZipString(str string) string {

	newStr := sortString(str)
	strLen := len(newStr)
	toMap := make(map[string]int)
	var zipStr string

	for _, s := range newStr {
		toMap[string(s)]++
	}

	for val, _ := range toMap {
		zipStr += val + strconv.Itoa(toMap[val])
	}

	if len(zipStr) > strLen {
		return zipStr
	} else {
		return str
	}
}

func ZipString2(str string) string {

	newStr := sortString(str)

	var letter string
	var count int
	var zipStr []string
	//a2e1l2p2
	//a2e1l2p2
	//e1l2p2z1a2
	for i, s := range newStr {
		toStr := string(s)
		if letter != toStr || i+1 >= len(newStr) {

			if (letter != "" && letter != toStr) || i+1 >= len(newStr) {
				if i+1 == len(newStr) {
					count++
				}
				zipStr = append(zipStr, letter, strconv.Itoa(count))
			}
			letter = toStr
			count = 0
		}
		count++

	}

	if len(zipStr) > len(newStr) {
		return strings.Join(zipStr, "")
	} else {
		return str
	}
}
