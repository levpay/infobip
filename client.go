package infobip

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	//SingleMessagePath for sending a single message
	SingleMessagePath = "sms/1/text/single"
)

// HTTPInterface helps Infobip tests
type HTTPInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client manages requests to Infobip
type Client struct {
	BaseURL    string
	Username   string
	Password   string
	HTTPClient HTTPInterface
}

// ClientWithBasicAuth returns a pointer to infobip.Client with Infobip funcs
func ClientWithBasicAuth(username, password string) *Client {
	return &Client{
		BaseURL:    "https://api.infobip.com/",
		Username:   username,
		Password:   password,
		HTTPClient: &http.Client{},
	}
}

// SingleMessage sends one message to one recipient
func (c Client) SingleMessage(m Message) (r Response, err error) {
	if err = m.Validate(); err != nil {
		return
	}
	b, err := json.Marshal(m)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+SingleMessagePath, bytes.NewBuffer(b))
	if err != nil {
		return
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	return
}
