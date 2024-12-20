package encode_decode_tiny_url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	obj := Constructor()
	url := obj.encode("https://leetcode.com/problems/design-tinyurl")
	ans := obj.decode(url)
	assert.Equal(t, "https://leetcode.com/problems/design-tinyurl", ans)
}
