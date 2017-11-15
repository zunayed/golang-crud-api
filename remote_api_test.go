package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPriceLookup(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintln(w, "{\"ID\": \"10\", \"Price\": 5}")
	}))

	defer ts.Close()

	client := createClient(http.DefaultClient, ts.URL)
	_, err := client.LookupPrice("10")

	if err != nil {
		t.Errorf("LookupPrice() failed with error: %v", err)
		t.FailNow()
	}

}
