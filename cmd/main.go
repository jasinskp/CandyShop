package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jasinskp/CandyShop/pkg/http/rest"
)

func main() {
	fmt.Println("Server starting on port 8080")

	router := rest.InintHandlers()

	log.Fatal(http.ListenAndServe(":8080", router))
}
