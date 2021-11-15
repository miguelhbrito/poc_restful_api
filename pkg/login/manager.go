package login

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/storage"
)

type Manager struct {
	loginManager storage.LoginPostgres
}

func (m Manager) LoginIntoSystem(mctx mcontext.Context, l entity.LoginEntity) (string, error) {
	expectedPassword, ok := users[l.Cpf]
	if !ok || expectedPassword != l.Secret {
		return "", errors.New("User is not authorized")
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
		return "", err
	}

	if err := m.loginManager.SaveLogin(mctx, entity.LoginEntity{Cpf: l.Cpf, Secret: tokenString}); err != nil {
		return "", err
	}

	return tokenString, err
}
