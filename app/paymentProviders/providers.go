package paymentproviders

import "fmt"

type PaymentProvider interface {
	GetPaymentLink(price float64) string
}

type PayPalPaymentProvider struct{}

func (provider *PayPalPaymentProvider) GetPaymentLink(price float64) string {
	return fmt.Sprintf("Payment link from %T. Price: %v", *provider, price)
}

type ApplePayPaymentProvider struct{}

func (provider *ApplePayPaymentProvider) GetPaymentLink(price float64) string {
	return fmt.Sprintf("Payment link from %T. Price: %v", *provider, price)
}

type GooglePayPaymentProvider struct{}

func (provider *GooglePayPaymentProvider) GetPaymentLink(price float64) string {
	return fmt.Sprintf("Payment link from %T. Price: %v", *provider, price)
}
