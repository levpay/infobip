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

func TestForResponseMessages(t *testing.T) {
	tests := []struct {
		reference string
		body      string
		data      Response
	}{
		{
			reference: "#1",
			body:      `{}`,
			data:      Response{},
		},
		{
			reference: "#2",
			body:      `{"messages": []}`,
			data: Response{
				Messages: []ResponseMessage{},
			},
		},
		{
			reference: "#3",
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
			reference: "#4",
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
			reference: "#5",
			body:      `{"messages": [{"to": "41793026727", "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
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
			reference: "#6",
			body:      `{"messages": [{"to": "41793026727", "status": {"id": 0}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
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
			reference: "#7",
			body:      `{"messages": [{"to": "41793026727", "status": {"id": 0, "groupId": 0}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
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
			reference: "#8",
			body:      `{"messages": [{"to": "41793026727", "status": {"id": 0, "groupId": 0, "groupName": "ACCEPTED"}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
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
			reference: "#9",
			body:      `{"messages": [{"to": "41793026727", "status": {"id": 0, "groupId": 0, "groupName": "ACCEPTED", "name": "MESSAGE_ACCEPTED"}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
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
			reference: "#10",
			body:      `{"messages": [{"to": "41793026727", "status": {"id": 0, "groupId": 0, "groupName": "ACCEPTED", "name": "MESSAGE_ACCEPTED", "description": "Message accepted"}, "smsCount": 1, "messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"}]}`,
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
	for _, test := range tests {
		test := test
		t.Run(test.reference, func(t *testing.T) {
			client.HTTPClient = &HTTPClientMock{
				Response: &http.Response{
					Body: ioutil.NopCloser(bytes.NewBufferString(test.body)),
				},
			}
			r, err := client.SingleMessage(message)
			if err != nil {
				t.Errorf("Error: unexpected error was returned (%s)", err)
			}
			if !reflect.DeepEqual(r, test.data) {
				t.Errorf("expected '%v', got '%v'", test.data, r)
			}
		})
	}
}

func TestForSingleMessageError(t *testing.T) {
	tests := []struct {
		reference string
		message   Message
		err       error
	}{
		{
			reference: "#1",
			message:   Message{},
			err:       ErrForFromNonAlphanumeric,
		},
	}

	client := ClientWithBasicAuth("foo", "bar")
	client.HTTPClient = &HTTPClientMock{
		Response: &http.Response{
			Body: ioutil.NopCloser(nil),
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.reference, func(t *testing.T) {
			if _, err := client.SingleMessage(test.message); err != test.err {
				t.Errorf("expected '%v', got '%v'", test.err, err)
			}
		})
	}
}

func TestForAdvancedMessageError(t *testing.T) {
	tests := []struct {
		reference string
		message   BulkMessage
		err       error
	}{
		{
			reference: "#1",
			message: BulkMessage{
				Messages: []Message{
					Message{},
				},
			},
			err: ErrForFromNonAlphanumeric,
		},
	}

	client := ClientWithBasicAuth("foo", "bar")
	client.HTTPClient = &HTTPClientMock{
		Response: &http.Response{
			Body: ioutil.NopCloser(nil),
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.reference, func(t *testing.T) {
			if _, err := client.AdvancedMessage(test.message); err != test.err {
				t.Errorf("expected '%v', got '%v'", test.err, err)
			}
		})
	}
}
