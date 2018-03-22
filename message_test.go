package infobip

import "testing"

func TestForMessageValidation(t *testing.T) {
	messages := []struct {
		from string
		to   string
		err  error
	}{
		{"", "", ErrForNonAlphanumeric},
		{"111111111111111", "", ErrForNonAlphanumeric},
		{"invalid1111111", "", ErrForAlphanumeric},
		{"valid", "111111111111111", ErrForNonAlphanumeric},
		{"442071838750", "14155552671", nil},
	}
	for _, message := range messages {
		m := Message{From: message.from, To: message.to}
		if err := m.Validate(); err != message.err {
			t.Errorf("Error: expected '%s', got '%s'", message.err, err)
		}
	}
}
