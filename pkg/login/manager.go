package login

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/api/response"
	"github.com/stone_assignment/pkg/auth"
	"github.com/stone_assignment/pkg/mcontext"
)

type manager struct {
	accountManager accounts.Account
	auth           auth.Auth
}

func NewManager(accountManager accounts.Account, auth auth.Auth) Login {
	return manager{
		accountManager: accountManager,
		auth:           auth,
	}
}

func (m manager) LoginIntoSystem(mctx mcontext.Context, l entity.LoginEntity) (response.LoginToken, error) {
	//Getting credentials from database
	lr, err := m.accountManager.GetByCpf(mctx, l.Cpf)
	if err != nil {
		return response.LoginToken{}, err
	}

	//Checking input secretHash with secretHash from database
	check := m.auth.CheckPasswordHash(l.Secret, lr.Secret)
	if !check {
		return response.LoginToken{}, errUserOrPassIncorrect
	}

	//Generation jwt token
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Cpf: l.Cpf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Signing jwt token with our key
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
