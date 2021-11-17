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
	loginManager storage.Account
}

func NewManager(loginManager storage.Account) Login {
	return Manager{
		loginManager: loginManager,
	}
}

func (m Manager) LoginIntoSystem(mctx mcontext.Context, l entity.LoginEntity) (response.LoginToken, error) {
	//Getting credentials from database
	lr, err := m.loginManager.GetCredentials(mctx, l.Cpf)
	if err != nil {
		return response.LoginToken{}, err
	}

	//Checking input secretHash with secretHash from database
	check := auth.CheckPasswordHash(l.Secret, lr.Secret)
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
