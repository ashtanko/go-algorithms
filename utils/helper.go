package utils

import (
	"fmt"
	"gotest.tools/v3/assert"
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
