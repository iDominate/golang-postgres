package models

type Response struct {
	StatusCode int16       `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
