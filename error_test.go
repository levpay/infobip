package infobip

import "testing"

func TestForErrorFunc(t *testing.T) {
	err := Error{Err: "error message"}.Error()
	expected := `{"error":"error message"}`
	if err != expected {
		t.Errorf("Error: expected '%s', got '%s'", expected, err)
	}
}
