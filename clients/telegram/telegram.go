package telegram

import (
	"net/http"
)

type Client struct {
	host     string
	basePass string
	client   http.Client
}
func New(host string, token string) Client{
	return Client{
		host: host,
		basePass: FuncPass(),
		client: http.Client{},
	}
}

func FuncPass(token string) string{
	return "bor" + token
}
