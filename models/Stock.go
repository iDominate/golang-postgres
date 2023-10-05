package models

type Stock struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}
