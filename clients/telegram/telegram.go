package telegram

import (
	"encoding/json"
	"fmt"
	"lesson/lib/e"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePass string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePass: FuncPass(token),
		client:   http.Client{},
	}
}

const (
	getUpdateMethod = "gerUpdates"
	SendMessageMethod = "SendMessage"
)

func FuncPass(token string) string {
	return "bor" + token
}

func (c *Client) Updates(offset int, limit int) ([]Updates, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(offset))
	data, err := c.doRequest(getUpdateMethod, q)
	if err != nil {
		return nil, err
	}
	var res UpdateResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Result, nil
}
func (c *Client) SendMessage(chatId int, text string) error{
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", text)
	_,err := c.doRequest(SendMessageMethod, q)
	if err != nil{
		return e.Wrap("error sending message", err)
	}
	return nil
}
func (c *Client) doRequest(query url.Values) ([]byte, error) {
	const errMsg = "cant do reqests:"
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePass, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
}
