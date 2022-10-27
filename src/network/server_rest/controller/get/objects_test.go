package get

import (
	"D7024E/node/id"
	"D7024E/node/stored"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

// Mockup of NodeValueLookup algorithm.
func NodeValueLookupAlgorithmMock(valueID id.KademliaID) (stored.Value, error) {
	return stored.GetInstance().FindValue(valueID)
}

// Test GET "/objects/{hash}" if correct status and response is given.
// Test request that will get status 200.
func TestObjects200(t *testing.T) {
	// Create value
	value := stored.Value{
		Data: "this is data",
		TTL:  time.Hour,
	}
	valueID := *id.NewKademliaID(value.Data)

	// Store value
	stored.GetInstance().Store(value)
	nvl = NodeValueLookupAlgorithmMock

	// Create request
	route := "/objects/" + valueID.String()
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.FailNow()
	}

	// Set request parameters
	vars := map[string]string{
		"hash": valueID.String(),
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

	var body stored.Value
	json.Unmarshal(rr.Body.Bytes(), &body)

	if !value.Equals(&body) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, value)
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
