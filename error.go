package infobip

import "fmt"

var (
	// ErrForNonAlphanumeric for invalid numbers
	ErrForNonAlphanumeric = Error{Err: "non-alphanumeric 'From' value must be between 3 and 14 numbers"}

	// ErrForAlphanumeric for invalid names
	ErrForAlphanumeric = Error{Err: "alphanumeric 'From' value must be between 3 and 13 characters"}
)

// Error for Infobip
type Error struct {
	Err string `json:"error,omitempty"`
}

// Error func to implements error interface
func (e Error) Error() string {
	return fmt.Sprintf(`{"error":"%v"}`, e.Err)
}
