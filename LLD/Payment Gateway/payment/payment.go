package payment

import "fmt"

type PaymentType string

const (
	CREDIT_CARD PaymentType = "Credit_Card"
	UPI         PaymentType = "Upi"
)

type IpaymentMehod interface {
	Pay(amount int) string
}

type CreditCard struct {
	CardName string
}

func (cc *CreditCard) Pay(amount int) string {
	return fmt.Sprintf("paid %d amount through credit card", amount)
}

func GetCreditCard(cardName string) IpaymentMehod {
	return &CreditCard{
		CardName: cardName,
	}
}

type Upi struct {
	UpiId string
}

func (upi *Upi) Pay(amount int) string {
	return fmt.Sprintf("paid %d amount  through upi", amount)
}

func GetUpi(upiId string) IpaymentMehod {
	return &Upi{
		UpiId: upiId,
	}
}

func PaymentMethodFactory(ptype PaymentType, name string) IpaymentMehod {
	var payment IpaymentMehod
	switch ptype {
	case CREDIT_CARD:
		payment = GetCreditCard(name)
	case UPI:
		payment = GetUpi(name)
	}
	return payment
}
