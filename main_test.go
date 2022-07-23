package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MaxDMT16/go-web-api/app"
	"github.com/MaxDMT16/go-web-api/app/constants"
	"github.com/MaxDMT16/go-web-api/app/models"
)

var a app.App

func TestMain(m *testing.M) {
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestCorrectPaymentLink(t *testing.T) {
	routes := []string{
		"/api/pay-pal/3123",
		"/api/apple-pay/3123",
		"/api/apple-pal/3123",
		"/api/pay-pal/3123",
		"/api/google-pay/3123",
		"/api/google-pay/3123",
	}

	for _, route := range routes {
		response := getPaymentLinkByRoute(route, t)

		if response.Link == constants.GetHeadwayLink {
			t.Error("Expected payment link")
		}
	}
}

func TestPaymentLinkInvalidIdFormat(t *testing.T) {
	routes := []string{
		"/api/pay-pal/31f23",
		"/api/apple-pay/dsfdsf23",
		"/api/pay-pal/31,23",
	}

	for _, route := range routes {
		response := getPaymentLinkByRoute(route, t)

		if response.Link != constants.GetHeadwayLink {
			t.Errorf("Route %v is invalid. Headway link is expected", route)
		}
	}
}

func getPaymentLinkByRoute(route string, t *testing.T) models.GetPaymentLinkResponse {
	req, _ := http.NewRequest("GET", route, nil)

	response := executeRequest(req)

	body := response.Body.Bytes()
	var getPaymentLinkResponse models.GetPaymentLinkResponse
	json.Unmarshal(body, &getPaymentLinkResponse)
	return getPaymentLinkResponse
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
