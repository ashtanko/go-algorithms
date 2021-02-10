package utils

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func Checkf(t *testing.T, comparison assert.BoolOrComparison, testCase interface{}) {
	assert.Check(t, comparison, fmt.Sprintf("%+v", testCase))
}
