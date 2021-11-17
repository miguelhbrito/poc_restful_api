package response

type Account struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Cpf       string  `json:"cpf"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"createdAt"`
}
