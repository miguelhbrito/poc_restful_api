package mhttp

import (
	"net/http"
)

type HttpHandler interface {
	Handler() http.HandlerFunc
}
