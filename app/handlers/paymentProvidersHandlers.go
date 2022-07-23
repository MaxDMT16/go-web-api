package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"getheadway/app/constants"
	"getheadway/app/db"
	"getheadway/app/models"
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
	response := models.GetPaymentLinkResponse{
		Link: constants.GetHeadwayLink,
	}

	subscriptionId, err := getSubscriptionIdFromPath(r)
	if err != nil {
		respondJson(w, http.StatusBadRequest, response)
		return
	}

	subscription, err := getSubscriptionFromDb(subscriptionId)
	if err != nil {
		respondJson(w, http.StatusBadRequest, response)
		return
	}

	paymentLink, err := paymentProvider.GetPaymentLink(subscription.Price)
	if err != nil {
		respondJson(w, http.StatusBadRequest, response)
		return
	}

	response.Link = paymentLink

	respondJson(w, http.StatusOK, response)
}

func respondJson(w http.ResponseWriter, statusCode int, content interface{}) {
	response, err := json.Marshal(content)

	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("some error occured"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func getSubscriptionIdFromPath(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}

func getSubscriptionFromDb(subscriptionId int) (models.Subscription, error) {
	repository := db.SubscriptionsRepository{}
	return repository.GetSubscriptionById(subscriptionId)
}
