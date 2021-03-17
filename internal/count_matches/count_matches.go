package count_matches

const (
	itemType  = iota
	itemColor = iota
	itemName  = iota
)

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ans := 0
	if len(items) == 0 || ruleKey == "" || ruleValue == "" {
		return ans
	}

	keys := map[string]int{
		"type":  itemType,
		"color": itemColor,
		"name":  itemName,
	}

	for _, row := range items {
		if row[keys[ruleKey]] == ruleValue {
			ans++
		}
	}

	return ans
}
