package infobip

import "regexp"

// Message contains the body request
type Message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

// Validate validates the body request values
func (m Message) Validate() (err error) {
	err = m.validateFromValue()
	if err != nil {
		return
	}
	err = m.validateToValue()
	return
}

func (m Message) validateFromValue() (err error) {
	if m.isNumeric(m.From) && !m.isValidRange(m.From, 3, 14) {
		err = ErrForNonAlphanumeric
		return
	}
	if !m.isValidRange(m.From, 3, 13) {
		err = ErrForAlphanumeric
		return
	}
	return
}

func (m Message) validateToValue() (err error) {
	if m.isNumeric(m.To) && !m.isValidRange(m.To, 3, 14) {
		err = ErrForNonAlphanumeric
		return
	}
	return
}

func (m Message) isNumeric(s string) bool {
	return regexp.MustCompile(`^[\d]*$`).MatchString(s)
}

func (m Message) isValidRange(s string, a, b int) bool {
	return len(s) > a && len(s) <= b
}
