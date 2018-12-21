package infobip

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

const (
	//SingleMessagePath for sending a single message
	SingleMessagePath = "sms/1/text/single"

	//AdvancedMessagePath for sending advanced messages
	AdvancedMessagePath = "sms/1/text/advanced"

	// ReportsPath for getting SMS reports.
	ReportsPath = "/sms/1/reports"
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

func (c Client) GetDeliveryReport(smsID string) (*SmsReportResponse, error) {
	res := SmsReportResponse{}
	err := c.doRequest("GET", c.BaseURL+ReportsPath+"?messageId="+smsID, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
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

func (c *Client) doRequest(method string, path string, payload io.Reader, result interface{}) error {
	req, err := http.NewRequest(method, path, payload)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("User-Agent", "go-infobip/0.1")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		resp.Body.Close()
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return errors.New(resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(result)

	return err
}
