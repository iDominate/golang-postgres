package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iDominate/golang-postgres/models"
	"github.com/iDominate/golang-postgres/repos"
	"github.com/iDominate/golang-postgres/utils"
)

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	repo := repos.GetStockRepo()
	stocks, err := repo.GetAllStocks()

	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, nil, err.Error())
	}

	utils.ReturnResponse(w, http.StatusOK, stocks, "Success")
}

func Getstock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadGateway, nil, err.Error())
		return
	}
	repo := repos.GetStockRepo()
	stock, err := repo.GetStock(int64(id))

	if err != nil {
		fmt.Println(err.Error())
		utils.ReturnResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	utils.ReturnResponse(w, http.StatusOK, stock, "Success")

}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	json.NewDecoder(r.Body).Decode(&stock)
	repo := repos.GetStockRepo()
	s, err := repo.CreateStock(stock)

	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	utils.ReturnResponse(w, http.StatusCreated, s, "Success")
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	jsonId := mux.Vars(r)
	var stock models.Stock
	json.NewDecoder(r.Body).Decode(&stock)
	fmt.Println(stock)
	id, err := strconv.Atoi(jsonId["id"])

	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	repo := repos.GetStockRepo()
	s, error1 := repo.UpdateStock(int64(id), stock)
	if error1 != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	utils.ReturnResponse(w, http.StatusOK, s, "Success")

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	error1 := repos.GetStockRepo().DeleteStock(int64(id))

	if error1 != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	utils.ReturnResponse(w, http.StatusAccepted, id, "Success")
}
