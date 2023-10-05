package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iDominate/golang-postgres/controllers"
)

func RegisterStocksRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/stocks/", controllers.GetAllStocks).Methods(http.MethodGet)
	router.HandleFunc("/api/stocks/{id}", controllers.Getstock).Methods(http.MethodGet)
	router.HandleFunc("/api/stocks/{id}", controllers.UpdateStock).Methods(http.MethodPut)
	router.HandleFunc("/api/stocks/", controllers.CreateStock).Methods(http.MethodPost)
	router.HandleFunc("/api/stocks/{id}", controllers.DeleteStock).Methods(http.MethodDelete)

	return router
}
