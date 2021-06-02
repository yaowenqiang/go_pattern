package creational

import (
    "testing"
    "strings"
)

func TestGetPaymentMethodCash(t *testing.T) {

    payment, err := GetPaymentMethod(Cash)
    if err != nil {
        t.Fatal("A payment method of type 'Cash' must exist!")
    }

    msg := payment.Pay(10.30)
    if !strings.Contains(msg, "payed using cash") {
        t.Error("The cash payment method message wasn't correct")
    }

    t.Log("LOG:", msg)
}

