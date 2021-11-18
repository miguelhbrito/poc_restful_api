package response

import "time"

type Transfer struct {
	Id              string    `json:"id"`
	AccountOriginId string    `json:"accountOriginId"`
	AccountDestId   string    `json:"accountDestId"`
	Ammount         float64   `json:"ammount"`
	CreatedAt       time.Time `json:"CreatedAt"`
}
