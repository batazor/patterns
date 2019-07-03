package factory_method

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash       = iota
	DebitCard  = iota
	CreditCard = iota
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("A payment method with ID %d not recognized\n", m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}
type CreditCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using credit card\n", amount)
}
