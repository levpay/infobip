package infobip

import (
	"github.com/shopspring/decimal"
)

// Amount is a thin wrapper around decimal to support json unmarshal.
type Amount struct {
	decimal.Decimal
}

// UnmarshalJSON is an implementation of Unmarshaler interface for the Amount type.
func (a *Amount) UnmarshalJSON(data []byte) error {
	v, err := decimal.NewFromString(string(data))
	if err != nil {
		return err
	}
	a.Decimal = v
	return nil
}

// SentSmsStatus indicates whether the message is successfully sent,
// not sent, delivered, not delivered, waiting for delivery or any other possible status.
type SentSmsStatus struct {
	GroupID     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Action      string `json:"action"`
}

// SentSmsPrice is a Sent SMS price info: amount and currency.
type SentSmsPrice struct {
	PricePerMessage Amount `json:"pricePerMessage"`
	Currency        string `json:"currency"`
}

// SentSmsError indicates whether the error occurred during the query execution.
type SentSmsError struct {
	GroupID     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanent   bool   `json:"permanent"`
}

// SentSmsReport is a message-specific delivery report.
type SentSmsReport struct {
	BulkID    string        `json:"bulkId"`
	To        string        `json:"to"`
	SentAt    string        `json:"sentAt"`
	DoneAt    string        `json:"doneAt"`
	Status    SentSmsStatus `json:"status"`
	SmsCount  int           `json:"smsCount"`
	MessageID string        `json:"messageId"`
	MccMnc    string        `json:"mccMnc"`
	Price     SentSmsPrice  `json:"price"`
	Error     SentSmsError  `json:"error"`
}

// SmsReportResponse contains a collection of reports, one per every message.
type SmsReportResponse struct {
	Results []SentSmsReport `json:"results"`
}
