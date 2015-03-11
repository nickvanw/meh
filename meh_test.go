package meh

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	want := "test-key-123"
	conf := WithKey(want)
	cli := NewClient(conf)
	if cli.key != want {
		t.Fatalf("wanted key: %q, got: %q", want, cli.key)
	}
}

func TestReq(t *testing.T) {
	tUrl := testServer()
	wUrl := WithURL(tUrl)
	client := NewClient(wUrl)
	d, err := client.Current()

	want := new(Current)
	json.Unmarshal(testData, want)

	if err != nil {
		t.Fatalf("Wanted no error, got: %v", err)
	}
	if !reflect.DeepEqual(want, d) {
		t.Fatalf("Wanted %#v, got: %#v", want, d)
	}
}

func testServer() string {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(testData)
	}))
	return s.URL
}
