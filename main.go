package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func TimeRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request processed by: %s.\n", time.Since(start))
	})
}

func main() {
	cfg := NewConfig()
	handler := NewEtaService(cfg)

	log.Println("starting server at :8082")
	handler.WarmUp()

	err := http.ListenAndServe(cfg.Port, TimeRequestMiddleware(handler))
	if err != nil {
		fmt.Println("Listen error: ", err)
	}
}
