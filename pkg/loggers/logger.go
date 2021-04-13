package loggers

import (
	"log"
	"net/http"
)

// Wraps the response
func AffixLogger(h http.Handler, before string, after string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(before)
		h.ServeHTTP(w, r) // Response
		log.Println(after)
	})
}
