package infobip

type SearchNumberParmas struct {
	Number       string `json:"capabilities,omitempty" url:"capabilities,omitempty"`
	Capabilities string `json:"capabilities,omitempty" url:"capabilities,omitempty"`
	Country      string `json:"country,omitempty" url:"country,omitempty"`
	Limit        int    `json:"limit,omitempty" url:"limit,omitempty"`
	Page         int    `page:"limit,omitempty" url:"page,omitempty"`
}

type SearchNumberResponse struct {
	NumberCount int      `json:"numberCount,omitempty"`
	Numbers     []Number `json:"numbers,omitempty"`
}

type Number struct {
	NumberKey    string   `json:"numberKey,omitempty"`
	Number       string   `json:"number,omitempty"`
	Country      string   `json:"country,omitempty"`
	Type         string   `json:"type,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
}
