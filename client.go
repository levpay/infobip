package infobip

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	//SingleMessagePath for sending a single message
	SingleMessagePath = "sms/1/text/single"

	//AdvancedMessagePath for sending advanced messages
	AdvancedMessagePath = "sms/1/text/advanced"

	// AvailableNumberPath for searching number
	AvailableNumberPath = "numbers/1/numbers/available"
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

// SearchNumber return a list of available number
func (c Client) SearchNumber(parmas SearchNumberParmas) (*SearchNumberResponse, error) {
	v, err := query.Values(parmas)
	if err != nil {
		return nil, err
	}

	path := AvailableNumberPath + "?" + v.Encode()
	resp := &SearchNumberResponse{}

	if err := c.defaultGetRequest(path, resp); err != nil {
		return nil, err
	}

	return resp, nil
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

func (c Client) defaultGetRequest(path string, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}
