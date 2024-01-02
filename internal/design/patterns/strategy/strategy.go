package strategy

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

type Cash struct{}

type Bank struct{}

type PaymentStrategy interface {
	Pay(ctx *PaymentContext) string
}

func NewPayment(name, cardID string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardID,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() string {
	return p.strategy.Pay(p.context)
}

func (*Cash) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

func (*Bank) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)
}
