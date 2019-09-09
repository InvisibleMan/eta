package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := NewCmdConfig()
	handler := NewEtaService(NewModel(NewSwaggerApi(cfg.Endpoint)))

	log.Printf("Starting server at '%s'\n", cfg.Port)
	log.Printf("Use endpoint:  '%s'\n", cfg.Endpoint)

	handler.WarmUp()

	err := http.ListenAndServe(cfg.Port, TimeRequestMiddleware(handler))
	if err != nil {
		log.Println("Listen error: ", err)
	}
}
