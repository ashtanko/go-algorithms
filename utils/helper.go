package utils

import (
	"fmt"
	"gotest.tools/v3/assert"
	"math/rand"
	"testing"
)

func Checkf(t *testing.T, comparison assert.BoolOrComparison, testCase interface{}) {
	assert.Check(t, comparison, fmt.Sprintf("%+v", testCase))
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func GenerateRandArray(n int) []float64 {
	arr := make([]float64, n)
	for i := 0; i < cap(arr); i++ {
		arr[i] = rand.Float64() * float64(n)
	}
	return arr
}

func SameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

func SameString2DSlice(x, y [][]string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		for _, __x := range _x {
			// 0 value for int is 0, so just increment a counter for the string
			diff[__x]++
		}
	}
	for _, _y := range y {
		for _, __y := range _y {
			// If the string _y is not in diff bail out early
			if _, ok := diff[__y]; !ok {
				return false
			}
			diff[__y] -= 1
			if diff[__y] == 0 {
				delete(diff, __y)
			}
		}
	}
	return len(diff) == 0
}
