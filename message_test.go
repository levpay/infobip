package infobip

import "testing"

func TestForSingleMessageValidation(t *testing.T) {
	tests := []struct {
		reference string
		from      string
		to        string
		err       error
	}{
		{
			reference: "#1",
			from:      "",
			to:        "",
			err:       ErrForFromNonAlphanumeric,
		},
		{
			reference: "#2",
			from:      "111111111111111",
			to:        "",
			err:       ErrForFromNonAlphanumeric,
		},
		{
			reference: "#3",
			from:      "invalid1111111",
			to:        "",
			err:       ErrForFromAlphanumeric,
		},
		{
			reference: "#4",
			from:      "valid",
			to:        "111111111111111",
			err:       ErrForToNonAlphanumeric,
		},
		{
			reference: "#5",
			from:      "442071838750",
			to:        "14155552671",
			err:       nil,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.reference, func(t *testing.T) {
			m := Message{
				From: test.from,
				To:   test.to,
			}
			if err := m.Validate(); err != test.err {
				t.Errorf("Error: expected '%s', got '%s'", test.err, err)
			}
		})
	}
}

func TestForAdvancedMessageValidation(t *testing.T) {
	tests := []struct {
		reference string
		from      string
		to        []Destination
		err       error
	}{
		{
			reference: "#1",
			to:        []Destination{},
			err:       ErrForFromNonAlphanumeric,
		},
		{
			reference: "#2",
			from:      "111111111111111",
			to:        []Destination{},
			err:       ErrForFromNonAlphanumeric,
		},
		{
			reference: "#3",
			from:      "invalid1111111",
			to:        []Destination{},
			err:       ErrForFromAlphanumeric,
		},
		{
			reference: "#4",
			from:      "valid",
			to: []Destination{
				Destination{To: "111111111111111"},
			},
			err: ErrForDestinationNonAlphanumeric,
		},
		{
			reference: "#5",
			from:      "442071838750",
			to: []Destination{
				Destination{To: "14155552671"},
			},
			err: nil,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.reference, func(t *testing.T) {
			b := BulkMessage{
				Messages: []Message{
					Message{
						From:         test.from,
						Destinations: test.to,
					},
				},
			}
			if err := b.Validate(); err != test.err {
				t.Errorf("Error: expected '%s', got '%s'", test.err, err)
			}
		})
	}
}
