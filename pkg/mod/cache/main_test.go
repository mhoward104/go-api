package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//make a test call to the api to check return status, and contents.
func TestGetInfo(t *testing.T) {
	s := newServerStatus()

	request, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(s.getInfo)
	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("returned wrong status got: %v, expected: %v", status, http.StatusOK)
	}

	if e, err := json.Marshal(s); err == nil {
		if strings.TrimSpace(recorder.Body.String()) != strings.TrimSpace(string(e)) {
			t.Errorf("unexpected body returned:\ngot: %v \nwanted: %v", recorder.Body.String(), string(e))
		}
	}
}
