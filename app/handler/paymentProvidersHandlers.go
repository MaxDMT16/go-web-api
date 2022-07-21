package handler

import (
	"encoding/json"
	"net/http"

	"getheadway/app/db"
	"getheadway/app/model"
	"getheadway/app/paymentproviders"

	"github.com/gorilla/mux"
)

func GetPayPalPaymentLink(w http.ResponseWriter, r *http.Request) {
	getPaymentLink(w, r, &paymentproviders.PayPalPaymentProvider{})
}

func GetApplePayPaymentLink(w http.ResponseWriter, r *http.Request) {
	getPaymentLink(w, r, &paymentproviders.ApplePayPaymentProvider{})
}

func GetGooglePayPaymentLink(w http.ResponseWriter, r *http.Request) {
	getPaymentLink(w, r, &paymentproviders.GooglePayPaymentProvider{})
}

func getPaymentLink(w http.ResponseWriter, r *http.Request, paymentProvider paymentproviders.PaymentProvider) {
	subscriptionId := getSubscriptionIdFromPath(r)

	subscription := getSubscriptionFromDb(subscriptionId)

	paymentLink := paymentProvider.GetPaymentLink(subscription.Price)

	respondJson(w, paymentLink)
}

func respondJson(w http.ResponseWriter, content interface{}) {
	response, err := json.Marshal(content)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("some error occured"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func getSubscriptionIdFromPath(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}

func getSubscriptionFromDb(subscriptionId string) model.Subscription {
	db := db.DbContext{}

	return db.GetSubscriptionById(subscriptionId)
}
