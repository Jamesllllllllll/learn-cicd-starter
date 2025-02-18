package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{
		"Authorization": {"ApiKey 12345"},
	}
	got, _ := GetAPIKey(headers)
	want := "12345"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	headers2 := http.Header{
		"Authorization": {""},
	}
	_, got2 := GetAPIKey(headers2)
	want2 := errors.New("no authorization header included")
	if !reflect.DeepEqual(want2, got2) {
		t.Fatalf("expected: %v, got: %v", want2, got2)
	}
}
