package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Arrange
	testCases := []struct {
		key         string
		value       string
		expect      string
		expectedErr string
	}{
		{
			key:         "Authorization",
			value:       "ApiKey xxxxxxxx",
			expect:      "xxxxxxxx",
		},
		{
			key:         "Authorization",
			value:       "",
			expect:      "",
			expectedErr: "no authorization header included",
		},
		{
			key:         "Authorization",
			value:       "Bearer xxxxxxxx",
			expect:      "",
			expectedErr: "malformed authorization header",
		},
	}

	// Act
	for _, test := range testCases[:] {
		header := http.Header{}
		header.Add(test.key, test.value)

		output, err := GetAPIKey(header)
		expect := test.expect

		// Assert
		if !reflect.DeepEqual(expect, output) {
			t.Errorf("expected value: %v, Output value: %v", expect, output)
		}

		if err != nil {
			if reflect.DeepEqual(test.expectedErr, err) {
				t.Errorf("Expected error: %v, Output error: %v", test.expectedErr, err)
			}
		}
	}
}
