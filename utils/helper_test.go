package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestMin(t *testing.T) {
	testCases := []struct {
		x, y     int
		expected int
	}{
		{
			x:        3,
			y:        7,
			expected: 3,
		},
		{
			x:        -1,
			y:        -7,
			expected: -7,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("x: %d y: %d", tc.x, tc.y)
		t.Run(name, func(t *testing.T) {
			actual := Min(tc.x, tc.y)
			Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		x, y     int
		expected int
	}{
		{
			x:        3,
			y:        7,
			expected: 7,
		},
		{
			x:        -1,
			y:        -7,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("x: %d y: %d max=> %d", tc.x, tc.y, tc.expected)
		t.Run(name, func(t *testing.T) {
			actual := Max(tc.x, tc.y)
			Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}

func TestGenerateRandArray(t *testing.T) {
	arr := GenerateRandArray(100)
	assert.NotNil(t, arr)
	assert.Equal(t, 100, len(arr))
}
