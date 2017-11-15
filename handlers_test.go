package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/product/45", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	t.Logf("%v", m)

	if m["error"] != "Product not found" {
		t.Errorf("Expected the error. Got '%s'", m["error"])
	}
}
