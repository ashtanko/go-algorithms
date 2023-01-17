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

func TestSameStringSlice(t *testing.T) {

	t.Run("when empty slice, should be true", func(t *testing.T) {
		actual := SameStringSlice([]string{}, []string{})
		assert.True(t, actual)
	})

	t.Run("when slice with empty string, should be true", func(t *testing.T) {
		actual := SameStringSlice([]string{""}, []string{""})
		assert.True(t, actual)
	})

	t.Run("when slice with diff empty string, should be false", func(t *testing.T) {
		actual := SameStringSlice([]string{""}, []string{})
		assert.False(t, actual)
	})

	t.Run("when diff order, should be true", func(t *testing.T) {
		actual := SameStringSlice([]string{"1", "2"}, []string{"2", "1"})
		assert.True(t, actual)
	})

	t.Run("when diff value, should be false", func(t *testing.T) {
		actual := SameStringSlice([]string{"1", "2"}, []string{"2", "3"})
		assert.False(t, actual)
	})
}

func TestSameString2DSlice(t *testing.T) {
	t.Run("when empty, should be true", func(t *testing.T) {
		actual := SameString2DSlice([][]string{}, [][]string{})
		assert.True(t, actual)
	})

	t.Run("when diff order, should be true", func(t *testing.T) {
		actual := SameString2DSlice([][]string{{"1", "2"}}, [][]string{{"2", "1"}})
		assert.True(t, actual)
	})

	t.Run("when slice with empty string, should be true", func(t *testing.T) {
		actual := SameString2DSlice([][]string{{""}}, [][]string{{""}})
		assert.True(t, actual)
	})

	t.Run("when slice with diff empty string, should be false", func(t *testing.T) {
		actual := SameString2DSlice([][]string{{""}}, [][]string{})
		assert.False(t, actual)
	})

	t.Run("when diff value, should be false", func(t *testing.T) {
		actual := SameString2DSlice([][]string{{"1", "2"}}, [][]string{{"2", "3"}})
		assert.False(t, actual)
	})

	t.Run("when diff 2d order, should be true", func(t *testing.T) {
		actual := SameString2DSlice(
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
			[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		)
		assert.True(t, actual)
	})
}
