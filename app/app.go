package app

import (
	"log"
	"net/http"

	"getheadway/app/handler"
	"getheadway/app/middleware"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	a.initializeRoutes()

	a.Router.Use(middleware.EnrichWithHeadwayHeaders)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/pay-pal/{id}", handler.GetPayPalPaymentLink).Methods("GET")
	a.Router.HandleFunc("/api/apple-pay/{id}", handler.GetApplePayPaymentLink).Methods("GET")
	a.Router.HandleFunc("/api/google-pay/{id}", handler.GetGooglePayPaymentLink).Methods("GET")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// func (a *App) ErrorHandler(next http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				log.Printf("panic: %+v", err)
// 				http.Error(w, http.StatusText(500), 500)
// 			}
// 	}()

// 	next.ServeHTTP(w, r)
// 	}

// 	return http.HandlerFunc(fn)
// }
