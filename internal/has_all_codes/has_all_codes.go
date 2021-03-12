package has_all_codes

func hasAllCodes(s string, k int) bool {
	need := uint(1 << k)
	got := make([]bool, need)
	allOne := need - 1
	hashVal := uint(0)

	for i := 0; i < len(s); i++ {
		b1 := hashVal << 1
		b2 := b1 & allOne
		b3 := uint(s[i] - '0')
		b4 := b3 | b2
		hashVal = b4

		if i >= k-1 && !got[hashVal] {
			got[hashVal] = true
			need--
			if need == 0 {
				return true
			}
		}
	}

	return false
}
