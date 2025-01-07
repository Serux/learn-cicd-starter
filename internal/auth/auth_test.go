package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	got, _ := GetAPIKey(http.Header{"Authorization": []string{"ApiKey 1234"}})
	want := "12345"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("BAD")
	}
}
