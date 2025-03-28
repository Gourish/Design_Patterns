package payment

type IPaymentGateway interface {
	CreateCreditCard(string) IpaymentMehod
	CreateUPI(string) IpaymentMehod
}

type PayPal struct{}

func (pp *PayPal) CreateCreditCard(cardName string) IpaymentMehod {
	return GetCreditCard(cardName)
}

func (pp *PayPal) CreateUPI(upiId string) IpaymentMehod {
	return GetUpi(upiId)
}

type Stripe struct{}

func (st *Stripe) CreateCreditCard(cardName string) IpaymentMehod {
	return GetCreditCard(cardName)
}

func (st *Stripe) CreateUPI(upiId string) IpaymentMehod {
	return GetUpi(upiId)
}
