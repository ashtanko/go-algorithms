package count_matches

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestCountMatches(t *testing.T) {
	testCases := []struct {
		items     [][]string
		ruleKey   string
		ruleValue string
		expected  int
	}{
		{
			items: [][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "lenovo"},
				{"phone", "gold", "iphone"},
			},
			ruleKey:   "color",
			ruleValue: "silver",
			expected:  1,
		},
		{
			items: [][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "phone"},
				{"phone", "gold", "iphone"},
			},
			ruleKey:   "type",
			ruleValue: "phone",
			expected:  2,
		},
		{
			items:     [][]string{},
			ruleKey:   "",
			ruleValue: "",
			expected:  0,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%s", tc.items)
		t.Run(name, func(t *testing.T) {
			actual := countMatches(tc.items, tc.ruleKey, tc.ruleValue)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
