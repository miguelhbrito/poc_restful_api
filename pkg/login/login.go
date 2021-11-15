package login

import (
	"github.com/golang-jwt/jwt"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

var JwtKey = []byte("my_secret_key")

var users = map[string]string{
	"11815458623": "mithrandir",
	"12345678910": "gandalf",
}

type Claims struct {
	Cpf string `json:"cpf"`
	jwt.StandardClaims
}

type LoginManager interface {
	LoginIntoSystem(mctx mcontext.Context, l entity.LoginEntity) (string, error)
}
