package infobip

// Response body response
type Response struct {
	BulkID   string            `json:"bulkId,omitempty"`
	Messages []ResponseMessage `json:"messages"`
}

// ResponseMessage ...
type ResponseMessage struct {
	ID       string         `json:"messageId"`
	To       string         `json:"to"`
	Status   ResponseStatus `json:"status"`
	SMSCount int            `json:"smsCount"`
}

// ResponseStatus ...
type ResponseStatus struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
