package infobip

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type HTTPClientMock struct {
	Response *http.Response
}

func (c *HTTPClientMock) Do(req *http.Request) (*http.Response, error) {
	return c.Response, nil
}

func TestForSingleResponseMessages(t *testing.T) {
	response := []struct {
		body string
		data Response
	}{
		{
			body: `{}`,
			data: Response{},
		},
		{
			body: `{"messages": []}`,
			data: Response{
				Messages: []ResponseMessage{},
			},
		},
		{
			body: `{
				"messages": [{"to": "41793026727"}]}`,
			data: Response{
				Messages: []ResponseMessage{
					{
						To: "41793026727",
					},
				},
			},
		},
		{
			body: `{
				"messages": [{"to": "41793026727", "smsCount": 1}]}`,
			data: Response{
				Messages: []ResponseMessage{
					{
						To:       "41793026727",
						SMSCount: 1,
					},
				},
			},
		},
		{
			body: `{"messages": [{"to": "41793026727", "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
			data: Response{
				Messages: []ResponseMessage{
					{
						To:       "41793026727",
						SMSCount: 1,
						ID:       "2250be2d4219-3af1-78856-aabe-1362af1edfd2",
					},
				},
			},
		},
		{
			body: `{"messages": [{"to": "41793026727", "status": {"id": 0}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
			data: Response{
				Messages: []ResponseMessage{
					{
						To: "41793026727",
						Status: ResponseStatus{
							ID: 0,
						},
						SMSCount: 1,
						ID:       "2250be2d4219-3af1-78856-aabe-1362af1edfd2",
					},
				},
			},
		},
		{
			body: `{"messages": [{"to": "41793026727", "status": {"id": 0, "groupId": 0}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
			data: Response{
				Messages: []ResponseMessage{
					{
						To: "41793026727",
						Status: ResponseStatus{
							ID:      0,
							GroupID: 0,
						},
						SMSCount: 1,
						ID:       "2250be2d4219-3af1-78856-aabe-1362af1edfd2",
					},
				},
			},
		},
		{
			body: `{"messages": [{"to": "41793026727", "status": {"id": 0, "groupId": 0, "groupName": "ACCEPTED"}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
			data: Response{
				Messages: []ResponseMessage{
					{
						To: "41793026727",
						Status: ResponseStatus{
							ID:        0,
							GroupID:   0,
							GroupName: "ACCEPTED",
						},
						SMSCount: 1,
						ID:       "2250be2d4219-3af1-78856-aabe-1362af1edfd2",
					},
				},
			},
		},
		{
			body: `{"messages": [{"to": "41793026727", "status": {"id": 0, "groupId": 0, "groupName": "ACCEPTED", "name": "MESSAGE_ACCEPTED"}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
			data: Response{
				Messages: []ResponseMessage{
					{
						To: "41793026727",
						Status: ResponseStatus{
							ID:        0,
							GroupID:   0,
							GroupName: "ACCEPTED",
							Name:      "MESSAGE_ACCEPTED",
						},
						SMSCount: 1,
						ID:       "2250be2d4219-3af1-78856-aabe-1362af1edfd2",
					},
				},
			},
		},
		{
			body: `{"messages": [{"to": "41793026727", "status": {"id": 0, "groupId": 0, "groupName": "ACCEPTED", "name": "MESSAGE_ACCEPTED", "description": "Message accepted"}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
			data: Response{
				Messages: []ResponseMessage{
					{
						To: "41793026727",
						Status: ResponseStatus{
							ID:          0,
							GroupID:     0,
							GroupName:   "ACCEPTED",
							Name:        "MESSAGE_ACCEPTED",
							Description: "Message accepted",
						},
						SMSCount: 1,
						ID:       "2250be2d4219-3af1-78856-aabe-1362af1edfd2",
					},
				},
			},
		},
	}

	var message = Message{
		From: "company",
		To:   "442071838750",
		Text: "Foo bar",
	}

	client := ClientWithBasicAuth("foo", "bar")
	for _, response := range response {
		client.HTTPClient = &HTTPClientMock{
			Response: &http.Response{
				Body: ioutil.NopCloser(bytes.NewBufferString(response.body)),
			},
		}
		r, err := client.SingleMessage(message)
		if err != nil {
			t.Errorf("Error: unexpected error was returned (%s)", err)
		}
		if !reflect.DeepEqual(r, response.data) {
			t.Errorf("expected '%v', got '%v'", response.data, r)
		}
	}
}
