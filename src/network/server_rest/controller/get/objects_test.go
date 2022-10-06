package get

import (
	"D7024E/node/id"
	"D7024E/node/stored"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// Test GET "/objects/{hash}" if correct status and response is given.
// Test request that will get status 200.
func TestObjects200(t *testing.T) {
	// Create value
	value := stored.Value{Data: "this is data"}
	value.ID = *id.NewKademliaID(value.Data)
	jsonValue, err := json.Marshal(value)
	if err != nil {
		t.FailNow()
	}

	// Store value
	stored.GetInstance().Store([]stored.Value{value})

	// Create request
	route := "/objects/" + value.ID.String()
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.FailNow()
	}

	// Set request parameters
	vars := map[string]string{
		"hash": value.ID.String(),
	}
	req = mux.SetURLVars(req, vars)

	// Test the request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Objects)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := string(jsonValue)
	body := strings.ReplaceAll(rr.Body.String(), "\n", "")
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, expected)
	}
}

// Test GET "/objects/{hash}" if correct status and response is given.
// Test request that will get status 400.
func TestObjects400(t *testing.T) {
	valueID := id.NewRandomKademliaID().String() + "a"
	// Create request
	route := "/objects/" + valueID
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.FailNow()
	}

	// Set request parameters
	vars := map[string]string{
		"hash": valueID,
	}
	req = mux.SetURLVars(req, vars)

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

// Test GET "/objects/{hash}" if correct status and response is given.
// Test request that will get status 404.
func TestObjects404(t *testing.T) {
	valueID := id.NewRandomKademliaID().String()

	// Create request
	route := "/objects/" + valueID
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.FailNow()
	}

	// Set request parameters
	vars := map[string]string{
		"hash": valueID,
	}
	req = mux.SetURLVars(req, vars)

	// Test the request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Objects)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	// Check the response body is what we expect.
	expected := ""
	body := strings.ReplaceAll(rr.Body.String(), "\n", "")
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, expected)
	}
}
