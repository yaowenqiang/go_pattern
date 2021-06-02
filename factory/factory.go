package creational

import (
    _ "errors"
    "fmt"
)

type PaymentMethod interface {
    Pay(amount float32) string
}

const (
    Cash      = 1
    DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
    //return nil, errors.New("Not implemented yet")
    if m == Cash {
        return &CashPM{}, nil
    } else if m == DebitCard {
        return &DebitCardPM{}, nil
    }
    return nil, nil
}

type CashPM struct{}
type DebitCardPM struct{}
func (c *CashPM) Pay(amount float32) string{
    return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string{
    return ""
}

