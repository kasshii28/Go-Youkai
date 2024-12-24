package middlewarepattern

import (
	"log"
	"net/http"
	"time"
)

func MyMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		d := time.Now().Sub(start).Milliseconds()
		log.Printf("end %s(%d ms)\n", start.Format(time.RFC3339), d)
	})
}