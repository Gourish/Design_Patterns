package main

import (
	"fmt"

	"example.payment_gateway.com/payment"
)

func main() {
	var factory payment.IPaymentGateway = &payment.PayPal{}

	paypalCredit := factory.CreateCreditCard("123-456-789")
	fmt.Print(paypalCredit.Pay(250))

	factory = &payment.Stripe{}

	stripeUpi := factory.CreateUPI("123@ybl")

	fmt.Print(stripeUpi.Pay(124))
}
