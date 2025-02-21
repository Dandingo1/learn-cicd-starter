package auth

import (
	"reflect"
	"testing"
)

// Happy Path
func TestGetAPIKey(t * testing.T) {
	// Arrange
	testCases := []map[string][]string {
		{
			"Authorization": {"ApiKey some_token"},
		},
		{
			"Authorization": {""},
		},
		{
			"Authorization": {"Bearer some_token_bearer"},
		},	
	}

	// Act
	got, err := GetAPIKey(testCases[0])
	want := "some_token"
	// Assert
	if (!reflect.DeepEqual(want, got)) {
		t.Fatalf("expected: %v, got: %v, err: %v", want, got, err)
	}

	// Act
	for _, tc := range testCases[1:] {
		got, err := GetAPIKey(tc)
		want := ""
		// Assert
		if (reflect.DeepEqual(want, got)) {
			t.Errorf("expected: %v, got: %v, err: %v", want, got, err)
		}
	}
	
}