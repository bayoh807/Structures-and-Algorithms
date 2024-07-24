package array

func OneEditAway(str1, str2 string) bool {

	toMap := make(map[string]int)
	var max, min string
	var stringLength int
	if res := len(str1) - len(str2); res > 1 || res < -1 {
		return false
	} else if res == -1 {
		max = str2
		min = str1
	} else if res == 1 {
		max = str1
		min = str2
	} else {
		max = str1
		min = str2
	}

	for _, s := range max {
		toMap[string(s)]++
	}
	stringLength = len(max)

	for _, s := range min {
		if _, res := toMap[string(s)]; res {
			stringLength--
		}
	}
	return stringLength == 1 || stringLength == 0
}
