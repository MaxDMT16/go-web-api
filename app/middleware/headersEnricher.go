package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func EnrichWithHeadwayHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enrichWithServerNameHeader(w)
		enrichWithResponseTime(next, w, r)
	})
}

func enrichWithResponseTime(next http.Handler, w http.ResponseWriter, r *http.Request) {
	startExecutionTime := time.Now()
	fmt.Printf("response time middleware: Start tracking\n")
	w.Header().Set("Trailer", "X-Response-Time")

	next.ServeHTTP(w, r)

	requesteExecutionTime := time.Since(startExecutionTime).Microseconds()
	fmt.Printf("response time middleware: Execution time: %v\n", requesteExecutionTime)
	w.Header().Set("X-Response-Time", fmt.Sprintf("%v", requesteExecutionTime))
}

func enrichWithServerNameHeader(w http.ResponseWriter) {
	hostname, _ := os.Hostname() //	TODO: handle error
	w.Header().Set("X-Server-Name", hostname)
}
