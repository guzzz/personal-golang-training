package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Load()

	fmt.Println("Server running - localhost:8000")

	r := router.Generate()

	apiPort := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(apiPort, r))
}
