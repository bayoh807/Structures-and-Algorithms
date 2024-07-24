package array

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
