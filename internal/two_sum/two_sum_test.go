package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("should work correctly on case [2,7,11,15]", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := twoSum(nums, target)
		expected := []int{0, 1}
		assert.Equal(t, result, expected)
	})

	t.Run("should work correctly on case [3,2,4]", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6
		result := twoSum(nums, target)
		expected := []int{1, 2}
		assert.Equal(t, result, expected)
	})

	t.Run("should work correctly on case []", func(t *testing.T) {
		var nums []int
		target := 0
		result := twoSum(nums, target)
		expected := []int{}
		assert.Equal(t, result, expected)
	})
}
