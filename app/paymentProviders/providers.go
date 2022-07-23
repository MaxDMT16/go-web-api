package paymentproviders

import (
	"fmt"

	"github.com/MaxDMT16/go-web-api/app/constants"
)

type PaymentProvider interface {
	GetPaymentLink(price float64) (string, error)
}

type PayPalPaymentProvider struct{}

func (provider *PayPalPaymentProvider) GetPaymentLink(price float64) (string, error) {
	return fmt.Sprintf(constants.PaymentLinkTemplate, "pay-pal", price), nil
}

type ApplePayPaymentProvider struct{}

func (provider *ApplePayPaymentProvider) GetPaymentLink(price float64) (string, error) {
	return fmt.Sprintf(constants.PaymentLinkTemplate, "apple-pay", price), nil
}

type GooglePayPaymentProvider struct{}

func (provider *GooglePayPaymentProvider) GetPaymentLink(price float64) (string, error) {
	return fmt.Sprintf(constants.PaymentLinkTemplate, "google-pay", price), nil
}
