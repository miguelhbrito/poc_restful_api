package response

import "time"

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"Cpf"`
	Balance   float64   `json:"Balance"`
	CreatedAt time.Time `json:"CreatedAt"`
}
