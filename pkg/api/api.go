package api

const (
	AuthorizationCtxKey Context = "authorization"
	CpfCtxKey           Context = "cpf"
)

type (
	Context string
	Cpf     string
)

func (c Cpf) String() string {
	return string(c)
}
