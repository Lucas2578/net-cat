package utils

func VerifyIsEmpty(str string) bool {
	isEmpty := true
	if len(str) <= 0 {
		return isEmpty
	}

	for _, char := range str {
		if char == ' ' {
			continue
		} else {
			isEmpty = false
			break
		}
	}

	return isEmpty
}

func VerifyIsTooLong(str string) bool {
	if len(str) > 16 {
		return true
	}

	return false
}
