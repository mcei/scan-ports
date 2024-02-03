package main

import (
	"testing"
)

// TestValidate calls validate, checking
// for a valid return value.
func TestValidate(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"negative", "", false},
		{"positive", "localhost", true},
	}

	errorFmt := "Expected %v, got %v on validating %s host."

	t.Parallel()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := validate(c.input)
			if got != c.want {
				t.Errorf(errorFmt, c.want, got, c.input)
			}
		})
	}
}
