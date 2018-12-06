package infobip

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	//SingleMessagePath for sending a single message
	SingleMessagePath = "sms/1/text/single"

	//AdvancedMessagePath for sending advanced messages
	AdvancedMessagePath = "sms/1/text/advanced"
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
		BaseURL:  "https://api.infobip.com/",
		Username: username,
		Password: password,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
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
	r, err = c.defaultRequest(b, SingleMessagePath)
	return
}

// AdvancedMessage sends messages to the recipients
func (c Client) AdvancedMessage(m BulkMessage) (r Response, err error) {
	if err = m.Validate(); err != nil {
		return
	}
	b, err := json.Marshal(m)
	if err != nil {
		return
	}
	r, err = c.defaultRequest(b, AdvancedMessagePath)
	return
}

func (c Client) defaultRequest(b []byte, path string) (r Response, err error) {
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+path, bytes.NewBuffer(b))
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
