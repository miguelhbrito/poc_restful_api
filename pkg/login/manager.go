package login

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/api/response"
	"github.com/stone_assignment/pkg/auth"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/storage"
)

type Manager struct {
	loginManager storage.AccountPostgres
}

func (m Manager) LoginIntoSystem(mctx mcontext.Context, l entity.LoginEntity) (response.LoginToken, error) {
	lr, err := m.loginManager.GetCredentials(mctx, l.Cpf)
	if err != nil {
		return response.LoginToken{}, err
	}

	check := auth.CheckPasswordHash(l.Secret, lr.Secret)
	if !check {
		return response.LoginToken{}, errUserOrPassIncorrect
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Cpf: l.Cpf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return response.LoginToken{}, err
	}

	tokenResponse := response.LoginToken{
		Token:   tokenString,
		ExpTime: expirationTime.Unix(),
	}

	return tokenResponse, err
}
