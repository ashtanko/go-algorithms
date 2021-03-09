package strobogrammatic

func isStrobogrammatic(num string) bool {
	left, right := 0, len(num)-1
	for left < right {
		if !check(num[left], num[right]) {
			return false
		}
		left++
		right--
	}

	if left == right {
		return checkOne(num[left])
	}
	return true
}

func check(b1, b2 byte) bool {
	switch b1 {
	case '0':
		return b2 == '0'
	case '1':
		return b2 == '1'
	case '6':
		return b2 == '9'
	case '8':
		return b2 == '8'
	case '9':
		return b2 == '6'
	}
	return false
}

func checkOne(b1 byte) bool {
	return b1 == '0' || b1 == '1' || b1 == '8'
}
