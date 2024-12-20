package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrategyPattern(t *testing.T) {
	t.Run("Pay by cash", func(t *testing.T) {
		payment := NewPayment("Ada", "1", 12, &Cash{})
		actual := payment.Pay()
		assert.Equal(t, "Pay $12 to Ada by cash", actual)
	})

	t.Run("Pay by bank", func(t *testing.T) {
		payment := NewPayment("Bob", "2", 88, &Bank{})
		actual := payment.Pay()
		assert.Equal(t, "Pay $88 to Bob by bank account 2", actual)
	})
}
