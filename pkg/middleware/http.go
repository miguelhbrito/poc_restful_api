package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/stone_assignment/pkg/api"
	"github.com/stone_assignment/pkg/login"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mctx := mcontext.NewFrom(r.Context())
		mlog.Debug(mctx).Msgf("Authorization middleware")

		if r.URL.Path == "/transfers" {
			mlog.Debug(mctx).Msgf("Authorization middleware checking token auth")
			tokenAuth := r.Header.Get("authorization")
			token, err := jwt.Parse(tokenAuth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return login.JwtKey, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				cpf := claims["cpf"]
				mctx = mcontext.WithValue(mctx, "props", claims)
				mctx = mcontext.WithValue(mctx, api.UsernameCtxKey, cpf)
				mctx = mcontext.WithValue(mctx, api.AuthorizationCtxKey, tokenAuth)
				next(w, r.WithContext(mctx))
			} else {
				fmt.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			next(w, r.WithContext(mctx))
		}
	}
}
