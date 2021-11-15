package entity

import "github.com/stone_assignment/pkg/api/response"

type LoginEntity struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret`
}

func (l LoginEntity) Response(token, expTime string) response.LoginToken {
	return response.LoginToken{
		Token:   token,
		ExpTime: expTime,
	}
}
