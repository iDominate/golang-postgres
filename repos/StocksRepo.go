package repos

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/iDominate/golang-postgres/models"
	"github.com/iDominate/golang-postgres/utils"

	_ "github.com/lib/pq"
)

var once sync.Once

var repo *StocksRepo

type StocksRepo struct {
	db *sql.DB
}

func GetStockRepo() *StocksRepo {
	once.Do(func() {
		repo = &StocksRepo{utils.GetConnection()}
	})

	return repo
}

func CloseConnection() {
	repo.db.Close()
}

func (s *StocksRepo) GetAllStocks() ([]models.Stock, error) {
	sqlSatement := "SELECT * FROM stocks"

	rows, err := s.db.Query(sqlSatement)

	if err != nil {
		log.Fatal(err)
		return make([]models.Stock, 0), errors.New(err.Error())
	}

	var stocks []models.Stock

	for rows.Next() {
		var stock models.Stock

		rows.Scan(&stock.Id, &stock.Name, &stock.Price, &stock.Company)

		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func (s *StocksRepo) GetStock(id int64) (models.Stock, error) {
	var stock models.Stock
	sqlStatement := "SELECT * FROM stocks WHERE sockid=$1"
	s.db.QueryRow(sqlStatement, id).Scan(&stock.Id, &stock.Name, &stock.Price, &stock.Company)

	return stock, nil
}

func (s *StocksRepo) CreateStock(stock models.Stock) (models.Stock, error) {
	sqlStatement := "INSERT INTO stocks(name, price, company) VALUES($1, $2, $3) RETURNING sockid;"

	err := s.db.QueryRow(sqlStatement, &stock.Name, &stock.Price, &stock.Company).Scan(&stock.Id)
	if err != nil {
		return models.Stock{}, err
	}

	return stock, nil
}

func (s *StocksRepo) UpdateStock(id int64, stock models.Stock) (models.Stock, error) {
	sqlStatement := "UPDATE stocks SET name = $1, price = $2, company = $3 WHERE sockid = $4 RETURNING sockid, name, price, company"
	err := s.db.QueryRow(sqlStatement, &stock.Name, &stock.Price, &stock.Company, id).Scan(&stock.Id, &stock.Name, &stock.Price, &stock.Company)
	if err != nil {
		fmt.Print(err.Error())
		return models.Stock{}, err
	}
	return stock, nil
}

func (s *StocksRepo) DeleteStock(id int64) error {
	sqlStatement := "DELETE FROM stocks WHERE sockid = $1"
	_, err := s.db.Exec(sqlStatement, id)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
