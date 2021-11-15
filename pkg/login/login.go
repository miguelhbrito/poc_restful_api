package login

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

var (
	JwtKey                 = []byte("my_secret_key")
	errUserOrPassIncorrect = errors.New("Username or Password is incorrect")
)

type Claims struct {
	Cpf string `json:"cpf"`
	jwt.StandardClaims
}

type LoginManager interface {
	LoginIntoSystem(mctx mcontext.Context, l entity.LoginEntity) (string, error)
}
