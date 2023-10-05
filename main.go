package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iDominate/golang-postgres/routes"
)

func main() {
	r := mux.NewRouter()

	r = routes.RegisterStocksRoutes(r)
	fmt.Println("Started server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))

}
