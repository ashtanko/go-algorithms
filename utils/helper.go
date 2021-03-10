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
