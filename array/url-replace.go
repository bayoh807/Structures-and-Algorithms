package array

func UrlReplace(str string, length int) string {

	var spaceCount, newLength int

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}

	newLength = length + spaceCount*2
	result := make([]byte, newLength)

	for i := length - 1; i >= 0; i-- {
		if str[i] == ' ' {
			result[newLength-1] = '%'
			result[newLength-2] = '2'
			result[newLength-3] = '0'
			newLength = newLength - 3
		} else {

			result[newLength-1] = str[i]
			newLength--
		}
	}

	return string(result)
}
