package creational

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("Payment method of type 'Cash' should exist")
	}

	msg := payment.Pay(20.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("Payment method of type 'DebitCard' should exist")
	}

	msg := payment.Pay(20.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message was not correct")
	}
	t.Log("LOG:", msg)
}
