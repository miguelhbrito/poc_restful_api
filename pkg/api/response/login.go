package response

type LoginToken struct {
	Token   string `json:"token"`
	ExpTime int64  `json:"expTime"`
}
