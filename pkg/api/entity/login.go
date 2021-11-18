package entity

type LoginEntity struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}
