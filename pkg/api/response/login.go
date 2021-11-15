package response

type LoginToken struct {
	Token   string `json:"token"`
	ExpTime string `json:"expTime"`
}
