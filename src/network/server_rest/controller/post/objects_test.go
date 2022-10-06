package post

import (
	"D7024E/node/id"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type message struct {
	Data string `json:"data"`
}

// Test POST "/objects" if correct status and response is given. Test request that will get status 200.
func TestObjectsSuccess(t *testing.T) {
	// Create request
	jsonMessage, err := json.Marshal(message{Data: "THIS IS THE DATA"})
	if err != nil {
		t.FailNow()
	}
	req, err := http.NewRequest("POST", "/objects", bytes.NewBuffer(jsonMessage))
	if err != nil {
		t.FailNow()
	}

	// Test the request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Objects)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	expected := `"/objects/` + id.NewKademliaID("THIS IS THE DATA").String() + `"`
	body := strings.ReplaceAll(rr.Body.String(), "\n", "")
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Test POST "/objects" if correct status and response is given. Test request that will get status 400.
func TestObjectsFail(t *testing.T) {
	// Create request
	jsonMessage, err := json.Marshal(message{})
	if err != nil {
		t.FailNow()
	}
	req, err := http.NewRequest("POST", "/objects", bytes.NewBuffer(jsonMessage))
	if err != nil {
		t.FailNow()
	}

	// Test the request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Objects)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	expected := ""
	body := strings.ReplaceAll(rr.Body.String(), "\n", "")
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, expected)
	}
}
