package array

func OneEditAway(str1, str2 string) bool {

	len1, len2 := len(str1), len(str2)

	if abs(len1-len2) > 1 {
		return false
	} else if len1 > len2 {
		str1, str2 = str2, str1
		len1, len2 = len2, len1
	}

	diffCount, i, j := 0, 0, 0

	for i < len1 && j < len2 {
		if str1[i] != str2[j] {
			diffCount++
			if diffCount > 1 {
				return false
			}
		}
		i++
		j++
	}
	return diffCount <= 1

}

func abs(val int) int {
	if val < 0 {
		val = -val
	}
	return val
}
