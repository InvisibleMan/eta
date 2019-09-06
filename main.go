package main

import (
	"fmt"
	"net/http"
)

func main() {
	cfg := NewConfig()
	handler := NewEtaService(cfg)

	fmt.Println("starting server at :8082")

	err := http.ListenAndServe(cfg.Port, handler)
	if err != nil {
		fmt.Println("Listen error: ", err)
	}
}
