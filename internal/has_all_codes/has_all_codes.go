package has_all_codes

import "math"

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}
	hash := map[string]interface{}{}
	for i := 0; i <= len(s)-k; i++ {
		hash[s[i:i+k]] = nil
	}
	total := int(math.Pow(2, float64(k)))
	return len(hash) == total
}
