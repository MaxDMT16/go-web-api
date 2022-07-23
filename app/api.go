package app

import (
	"log"
	"net/http"

	"github.com/MaxDMT16/go-web-api/app/handlers"
	"github.com/MaxDMT16/go-web-api/app/middlewares"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	a.initializeRoutes()

	a.Router.Use(middlewares.EnrichWithHeadwayHeaders)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/pay-pal/{id}", handlers.GetPayPalPaymentLink).Methods("GET")
	a.Router.HandleFunc("/api/apple-pay/{id}", handlers.GetApplePayPaymentLink).Methods("GET")
	a.Router.HandleFunc("/api/google-pay/{id}", handlers.GetGooglePayPaymentLink).Methods("GET")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
